.DEFAULT_GOAL := dev
SHELL := /bin/bash

.PHONY: help
help: ### Show this help message
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*?### "} /^[a-zA-Z_-]+:.*?### / {printf "  %-20s - %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: dev
dev: ### Run the operator in development mode
	@go run cmd/operator/main.go

.PHONY: lint
lint: ### Run all linters
	@pre-commit run -a