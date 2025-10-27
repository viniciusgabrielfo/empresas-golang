package table

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	minTableColumns = 4
	tableSeparator  = "---"
	ReadmeFileName  = "README.md"

	errFileNotFound  = "README.md not found"
	errTableNotFound = "could not find table in README.md"
	errFileOpen      = "failed to open file"
	errFileRead      = "failed to read file"
	errFileCreate    = "failed to create file"
	errFileWrite     = "failed to write file"
)

// Company represents a company entry in the README table
type Company struct {
	NameLink     string `json:"name_link"`
	Stackshare   string `json:"stackshare"`
	JobsLink     string `json:"jobs_link"`
	ContractType string `json:"contract_type"`
	OriginalLine string `json:"original_line"`
}

var (
	companyNameRegex = regexp.MustCompile(`\[([^\]]+)\]`)
	articles         = []string{"A ", "O ", "De ", "Da ", "Do ", "The "}
)

// normalizeName removes common articles from company names for consistent sorting
func NormalizeName(name string) string {
	normalized := strings.TrimSpace(name)

	for _, article := range articles {
		if strings.HasPrefix(normalized, article) {
			normalized = normalized[len(article):]
			break
		}
	}

	return strings.ToLower(normalized)
}

// extractCompanyName extracts the company name from markdown link format [Name](url)
func ExtractCompanyName(companyLink string) string {
	matches := companyNameRegex.FindStringSubmatch(companyLink)
	if len(matches) > 1 {
		return strings.TrimSpace(matches[1])
	}
	return strings.TrimSpace(companyLink)
}

func FindTableBoundaries(lines []string) (int, int, error) {
	tableStart := -1
	tableEnd := -1

	for i, line := range lines {
		if strings.Contains(line, "| Nome") && strings.Contains(line, "| Stackshare") {
			tableStart = i
		} else if tableStart != -1 && strings.TrimSpace(line) == "" {
			tableEnd = i
			break
		}
	}

	if tableStart == -1 {
		return -1, -1, errors.New(errTableNotFound)
	}

	if tableEnd == -1 {
		tableEnd = len(lines)
	}

	return tableStart, tableEnd, nil
}

func ExtractCompanies(lines []string, tableStart, tableEnd int) ([]Company, error) {
	companies := make([]Company, 0, tableEnd-tableStart-2)

	for i := tableStart + 2; i < tableEnd; i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}

		company := parseTableLine(line)
		if company != nil {
			companies = append(companies, *company)
		}
	}

	return companies, nil
}

// parseTableLine parses a markdown table line and returns a Company struct
func parseTableLine(line string) *Company {
	line = strings.TrimSpace(line)
	if !strings.HasPrefix(line, "|") || strings.Contains(line, tableSeparator) {
		return nil
	}

	// Remove leading and trailing pipes
	content := line[1 : len(line)-1]
	parts := strings.Split(content, "|")

	if len(parts) < minTableColumns {
		return nil
	}

	// Trim spaces from all parts
	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}

	return &Company{
		NameLink:     parts[0],
		Stackshare:   parts[1],
		JobsLink:     parts[2],
		ContractType: parts[3],
		OriginalLine: line,
	}
}

func RebuildFileContent(lines []string, companies []Company, tableStart, tableEnd int) []string {
	newLines := make([]string, 0, len(lines))

	newLines = append(newLines, lines[:tableStart]...)
	newLines = append(newLines, lines[tableStart])
	newLines = append(newLines, lines[tableStart+1])

	for _, company := range companies {
		newLines = append(newLines, company.OriginalLine)
	}

	newLines = append(newLines, lines[tableEnd:]...)

	return newLines
}

func WriteFile(filePath string, lines []string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("%s: %w", errFileCreate, err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, line := range lines {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return fmt.Errorf("%s: %w", errFileWrite, err)
		}
	}

	return nil
}
