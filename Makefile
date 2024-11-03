.DEFAULT_GOAL := dev
SHELL := /bin/bash

.PHONY: dev
dev:
	@go run cmd/operator/main.go

.PHONY: lint
lint: ### Run all linters
	@pre-commit run -a