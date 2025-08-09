# CRUSH.md for Kuq

This file provides essential guidelines and commands for agentic coding agents working in this repository.

## Build/Test/Lint Commands
- **Build:** `go build -o bin/kuq ./main.go`  (or `just build`)
- **Run:** `go run ./main.go`
- **Test (all):** `go test -v ./...`         (or `just test`)
- **Vet:** `go vet ./...`                    (included in `just test`)
- **Lint:** `golangci-lint run ./...`        (or `just lint`)
- **Single test:** `go test -v -run <TestName> ./...`
- **Format:** `gofmt -w .` (auto-run in CI via golangci-lint)

## Code Style Guidelines
- **Imports:**
  - Use grouped imports (`stdlib`, third-party, local) as per `gofmt`/`goimports`.
  - Prefer minimal imports and avoid unused ones (as enforced by CI linters).
- **Formatting:**
  - Use `gofmt`; no custom formatting rules. Whitespace is enforced by CI linter.
- **Types & Naming:**
  - Use CamelCase for types, structs, interfaces, and exported functions.
  - Use lowerCamelCase for unexported vars and functions.
  - Exported vars, funcs, and types require a clear doc comment per Golang convention.
  - Package names should be short and lower case, without underscores.
- **Error Handling:**
  - Handle errors explicitly. No silent ignores. Bubble up or log/fatal as appropriate.
  - Prefer errors.Is()/errors.As() for error type checks.
  - Wrap errors with context using `fmt.Errorf("describe: %w", err)`.
- **Linting:**
  - Follow rules enforced by the extensive `.golangci.yml` config—see that file for the complete linter list.
  - Revive and staticcheck are enabled; follow modern Go idioms.
- **Testing:**
  - Use Go's built-in `testing` package. Place tests alongside source, named `*_test.go`.
  - For parallel tests, use `t.Parallel()` where possible.
  - Use table-driven tests for all but trivial functions.
- **Other:**
  - No Cursor or Copilot rules present (checked as of 2025-08-09).
  - Update `.gitignore` to include `.crush/` if adding agent data.

_This file is auto-generated for coding agents. Last updated 2025-08-09._
