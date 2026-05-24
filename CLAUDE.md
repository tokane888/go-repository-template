# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This is a Go monorepo template for building microservices with shared packages. It follows [Standard Go Project Layout](https://github.com/golang-standards/project-layout) with pre-configured development tools.

## Project Structure

```text
.
├── services/           # Individual microservices
│   └── sample/        # Sample service
│       ├── cmd/sample/ # Application entry point (main.go)
│       └── internal/  # Private application code
├── pkg/               # Packages shared across services
│   └── logger/        # Shared logging package using slog
├── .devcontainer/     # VS Code DevContainer configuration
└── .github/           # GitHub Actions workflows
```

## Development Commands

### Running Services

```bash
# Run from the service directory (e.g., services/sample/)
go run cmd/sample/main.go
```

### Linting

```bash
# Run golangci-lint (from any Go module directory)
golangci-lint run

# Auto-fix some issues
golangci-lint run --fix
```

### Formatting

```bash
# Format Go code (handled by gofumpt via golangci-lint)
golangci-lint run --fix

# Format other files (JSON, Markdown, YAML, TOML)
dprint fmt

# Check formatting without changes
dprint check
```

### Module Management

```bash
# Update service dependencies
cd services/sample
go mod tidy

# Update all modules
find . -name go.mod -exec dirname {} \; | xargs -I {} sh -c 'cd {} && go mod tidy'
```

### Git Hooks

```bash
# Install git hooks (run from repository root)
lefthook install

# Run hooks manually
lefthook run pre-commit
lefthook run pre-push
```

## Architecture Decisions

1. **Monorepo structure**: Services and packages are managed in a single repository, making shared code management and consistent tooling easier.

2. **Internal packages**: Each service uses the `internal/` directory to prevent other services from importing private implementation details.

3. **Module boundaries**: Each service has its own `go.mod` file and uses `replace` directives for local packages during development.

4. **Configuration management**: Uses godotenv to load environment-specific configuration from `.env/.env.{ENV}` files.

5. **Structured logging**: All services use a shared logger package with slog, producing consistent JSON logs in production.

## Key Configuration Files

- `.golangci.yml`: Comprehensive linting rules including security checks, error handling, and style enforcement
- `dprint.json`: Formatting rules for non-Go files
- `.lefthook.yml`: Git hook configuration for automatic formatting and linting
- `.devcontainer/devcontainer.json`: VS Code development environment with all tools pre-installed

## Important TODOs

When using this template, address the following TODOs:

1. **Module name**: Update module paths in all `go.mod` files from `github.com/tokane888/go-repository-template` to your repository
2. **Import paths**: Update import statements in Go files to match the new module name
3. **Service name**: Rename the sample service and update configuration
4. **Environment variables**: Set appropriate values in `.env` files

## Testing Approach

The template does not include test files. When adding tests:

- Place unit tests alongside code files (e.g., `config_test.go`)
- Use `internal/testutil/` for test helpers
- Run tests from the service directory with `go test ./...`
- Use table-driven tests as the default style
- Name tests for a single function as `Test_validateConfig()` — `Test_` followed by the function name

## CI/CD Pipeline

GitHub Actions workflows handle:

- Running golangci-lint across all Go modules
- Checking code formatting with dprint
- Matrix builds across services
- Automatic PR checks

## Development Environment

DevContainer provides:

- Go 1.26.2
- golangci-lint
- dprint
- lefthook
- Git configuration
- Support for Claude Code and GitHub Copilot

## Source Editing Notes

- Do not delete comments without also removing or updating their corresponding source code
- When making fixes via GitHub issues and running `git commit`, use JST timezone

## golangci-lint Configuration

- Important: .golangci.yml uses v2 format

## Verification

- Run the following to format:
  - `gofumpt -w .`
- Run the following for spell check:
  - `cspell .`
- Confirm successful build
- Run `go test`
- Run `golangci-lint run ./...` in the directory containing the go.mod of the target process and confirm no warnings
- Public methods should have unit tests implemented, except for very simple ones
