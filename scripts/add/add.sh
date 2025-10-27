#!/bin/bash

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$(dirname "$SCRIPT_DIR")")"
README_PATH="$PROJECT_ROOT/README.md"

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}🏗️  Adicionando nova empresa ao README.md...${NC}"

if [ ! -f "$README_PATH" ]; then
    echo -e "${RED}❌ README.md não encontrado!${NC}"
    exit 1
fi

if ! command -v go &> /dev/null; then
    echo -e "${RED}❌ Go não encontrado!${NC}"
    echo -e "${YELLOW}💡 Instale Go para usar este script:${NC}"
    echo -e "${YELLOW}   https://go.dev/dl/${NC}"
    exit 1
fi

INPUT="$*"

if [ -z "$INPUT" ]; then
    echo -e "${RED}❌ Uso incorreto!${NC}"
    echo -e "${YELLOW}👉 make add \"Nome da empresa - link.site - link.stackshare or keep empty - link.vagas - tipo de contratação\"${NC}"
    exit 1
fi

echo -e "${YELLOW}🐹 Executando script para adicionar empresa...${NC}"
cd "$PROJECT_ROOT"

go run "$SCRIPT_DIR/add_company.go" "$INPUT"

echo -e "${GREEN}✅ Empresa adicionada com sucesso ao README.md!${NC}"