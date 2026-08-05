# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

A personal **Go learning playground**. It is not an application — it's a collection of small, independent example programs organized by topic under `basic_learning/`. The owner is learning Go in Spanish, so code comments, variable names, and explanations are in Spanish. Match that language and the simple, didactic style when adding or editing examples.

## Module layout

- Single Go module named `myProject` (see `go.mod`, Go 1.25.3).
- Internal imports therefore use the `myProject/...` prefix, e.g. `import "myProject/basic_learning/modules/saludar"`.
- Examples live under `basic_learning/<topic>/...` (punteros, slice, maps, structs, functions, errores, condicionales, etc.).
- Most example files are their own `package main` with a `func main()`, meant to be run one at a time — there is no single application entry point. The repo intentionally contains many `main` packages in separate directories.
- Some topics split logic into a sub-package (e.g. `functions/practica2/op_mat`, `modules/saludar`) that the topic's `main.go` imports via the `myProject/...` path — this is how multi-file/package examples are demonstrated.

## Commands

```bash
# Run a single example (the normal workflow — one file/folder at a time)
go run basic_learning/printing/imprimir.go

# Run an example that imports sibling packages: run the whole package dir, not one file
go run ./basic_learning/modules

# Build / vet everything
go build ./...
go vet ./...

# Format (run before committing)
gofmt -w .

# Tests (none exist yet; this is how you'd run them if added)
go test ./...
go test ./basic_learning/functions/...   # single package/dir
```

Note: `go run <file>.go` only works when the example is self-contained in that file. When an example imports a `myProject/...` sub-package, run the directory (`go run ./path/to/dir`) so all files in the package compile together.

## Environment

- `.env` is gitignored and loaded via `github.com/joho/godotenv` in examples that need it.
- The module pulls some heavy indirect deps (Google Cloud auth, protobuf, gRPC, gomail, logrus) from individual experiments; don't be surprised by the large `go.sum`.
