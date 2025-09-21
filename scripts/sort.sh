#!/bin/bash

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"
README_PATH="$PROJECT_ROOT/README.md"

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

echo -e "${BLUE}🔄 Ordenando empresas no README.md...${NC}"

if [ ! -f "$README_PATH" ]; then
    echo -e "${RED}❌ README.md não encontrado!${NC}"
    exit 1
fi

if ! command -v go &> /dev/null; then
    echo -e "${RED}❌ Go não encontrado!${NC}"
    echo -e "${YELLOW}💡 Instale Go para usar este script.${NC}"
    echo -e "${YELLOW}   https://go.dev/dl/${NC}"
    exit 1
fi

echo -e "${YELLOW}🐹 Executando script para ordenação...${NC}"
cd "$PROJECT_ROOT"
go run "$SCRIPT_DIR/sort_companies.go"