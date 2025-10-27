package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/viniciusgabrielfo/empresas-golang/scripts/table"
)

const (
	minTableColumns = 4
	tableSeparator  = "---"
	readmeFileName  = "README.md"

	errFileNotFound  = "README.md not found"
	errTableNotFound = "could not find table in README.md"
	errFileOpen      = "failed to open file"
	errFileRead      = "failed to read file"
	errFileCreate    = "failed to create file"
	errFileWrite     = "failed to write file"
	errArgInput      = "expected argument enclosed in quotes (\"\"), with 4 or 5 fields separated by hyphens ( - ). Example:\n make add \"Nome da empresa - link.site - link.stackshare or keep empty - link.vagas - tipo de contratação\""
)

type NewCompany struct {
	Name         string
	OficialLink  string
	Stackshare   string
	JobsLink     string
	ContractType string
}

func createCompany(args []string) (*NewCompany, error) {

	if len(args) <= 1 {
		return nil, fmt.Errorf("%s", errArgInput)
	}

	input := args[1]
	parts := strings.Split(input, " -")

	if len(parts) < 4 || len(parts) > 5 {
		return nil, fmt.Errorf("%s", errArgInput)
	}

	newCompany := &NewCompany{}

	if strings.Contains(parts[0], "http") {
		return nil, fmt.Errorf("first argument field must be the company name, not a link\n%s", errArgInput)
	}
	newCompany.Name = strings.TrimSpace(parts[0])

	if !strings.Contains(parts[1], "http") {
		return nil, fmt.Errorf("second argument field must be a link to the company's oficial site\n%s", errArgInput)
	}
	newCompany.OficialLink = strings.TrimSpace(parts[1])

	if !strings.Contains(parts[2], "http") {
		return nil, fmt.Errorf("third argument field must be a link to the company's stackshare profile, or must be omitted\n%s", errArgInput)
	}
	if !strings.Contains(parts[2], "stackshare.io") {
		newCompany.Stackshare = ""

		newCompany.JobsLink = strings.TrimSpace(parts[2])
	} else {
		newCompany.Stackshare = strings.TrimSpace(parts[2])
		if !strings.Contains(parts[3], "http") {
			return nil, fmt.Errorf("fourth argument field must be a link to the company's jobs page\n%s", errArgInput)
		}
		newCompany.JobsLink = strings.TrimSpace(parts[3])
	}

	contractField := strings.TrimSpace(parts[len(parts)-1])
	if !(contractField == "NACIONAL" || contractField == "INTERNATIONAL" || contractField == "NACIONAL/INTERNATIONAL") {
		return nil, fmt.Errorf("last argument field must be NACIONAL, INTERNATIONAL, or NATIONAL/INTERNATIONAL\n%s", errArgInput)
	}
	newCompany.ContractType = contractField

	return newCompany, nil
}

func addCompanyToReadme(readmePath string, newCompany *NewCompany) error {
	content, err := os.ReadFile(readmePath)
	if err != nil {
		return fmt.Errorf("%s: %w", errFileOpen, err)
	}

	lines := strings.Split(string(content), "\n")

	tableStart, tableEnd, err := table.FindTableBoundaries(lines)
	if err != nil {
		return err
	}

	companies, err := table.ExtractCompanies(lines, tableStart, tableEnd)
	if err != nil {
		return err
	}

	if err := duplicityCheck(companies, newCompany); err != nil {
		return err
	}

	companies = addNewCompany(companies, newCompany)

	newLines := table.RebuildFileContent(lines, companies, tableStart, tableEnd)

	if err := table.WriteFile(readmePath, newLines); err != nil {
		return err
	}

	return nil
}

func addNewCompany(companies []table.Company, newCompany *NewCompany) []table.Company {
	company := table.Company{}

	company.NameLink = fmt.Sprintf("[%s](%s)", newCompany.Name, newCompany.OficialLink)
	if newCompany.Stackshare != "" {
		company.Stackshare = fmt.Sprintf("[Clique aqui](%s)", newCompany.Stackshare)
	}
	company.JobsLink = fmt.Sprintf("[Clique aqui](%s)", newCompany.JobsLink)
	company.ContractType = newCompany.ContractType

	//TODO: format code table too, not just the markdown table
	company.OriginalLine = fmt.Sprintf("| %s | %s | %s | %s |", company.NameLink, company.Stackshare, company.JobsLink, company.ContractType)

	companies = append(companies, company)
	return companies
}

func duplicityCheck(companies []table.Company, newCompany *NewCompany) error {
	newName := table.NormalizeName(newCompany.Name)
	newName = strings.Join(strings.Fields(newName), "")

	for _, company := range companies {
		existingName := table.NormalizeName(table.ExtractCompanyName(company.NameLink))
		existingName = strings.Join(strings.Fields(existingName), "")

		if strings.Contains(existingName, newName) || strings.Contains(newName, existingName) {
			return fmt.Errorf("the company seems to already be in the table with this name: %s", table.ExtractCompanyName(company.NameLink))
		}
	}

	return nil
}

func main() {
	newCompany, err := createCompany(os.Args)
	if err != nil {
		log.Fatalf("❌ Parsing arguments: %v", err)
	}

	readmePath := table.ReadmeFileName

	if _, err := os.Stat(readmePath); os.IsNotExist(err) {
		log.Fatalf("❌ %s", errFileNotFound)
	}

	fmt.Println("🏗️ Adding company to README.md...")

	if err := addCompanyToReadme(readmePath, newCompany); err != nil {
		log.Fatalf("❌ Adding error: %v", err)
	}
}
