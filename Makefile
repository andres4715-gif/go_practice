# Makefile for the Go playground
# Usage: make <target>   (run `make help` to see everything)

# Current git branch (evaluated when used)
CURRENT_BRANCH := $(shell git branch --show-current)

.DEFAULT_GOAL := help

## help: Show this help
.PHONY: help
help:
	@echo "Available targets:"
	@grep -E '^## ' $(MAKEFILE_LIST) | sed 's/## /  /'

## fmt: Format all code with gofmt
.PHONY: fmt
fmt:
	go fmt ./...

## vet: Analyze all code with go vet
.PHONY: vet
vet:
	go vet ./...

## build: Compile all packages
.PHONY: build
build:
	go build ./...

## check: Format, analyze and build (all together)
.PHONY: check
check: fmt vet build

## push: git add + commit + push -> make push m="my message"
## Example: make push m="my commit message"
.PHONY: push
push:
	git add -A
	git commit -m "$(m)"
	git push origin $(CURRENT_BRANCH)

## tidy: Tidy up go.mod and go.sum
.PHONY: tidy
tidy:
	go mod tidy
