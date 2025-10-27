.PHONY: help sort test clean

GREEN=\033[0;32m
YELLOW=\033[1;33m
BLUE=\033[0;34m
NC=\033[0m

help: ## Mostra esta ajuda
	@echo "$(BLUE)📋 Comandos disponíveis:$(NC)"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  $(YELLOW)%-15s$(NC) %s\n", $$1, $$2}'

sort: ## Ordena as empresas na tabela do README.md
	@./scripts/sort.sh

add: ## Adiciona uma nova empresa à tabela do README.md > make add "Nome da empresa - link.site - link.stackshare or keep empty - link.vagas - tipo de contratação"
	@./scripts/add/add.sh $(filter-out $@,$(MAKECMDGOALS))
%:
	@:

.DEFAULT_GOAL := help
