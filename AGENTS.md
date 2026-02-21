# AGENTS.md

## Project Overview

Deepsearch-mockgen-cli is a Go CLI that scans a directory for Go interfaces and generates mocks using the `mockgen` binary (from `github.com/uber-go/mock`). It provides a TUI flow for selecting interfaces and confirming output paths, plus CLI flags for non-interactive runs.

Key components:

- [main.go](main.go) orchestrates CLI flags and TUI flow.
- [generator/](generator/) handles interface discovery, mock path computation, and `mockgen` invocation.
- [tui/](tui/) provides Bubble Tea-based prompts and progress UI.

## Setup Commands

- Install Go dependencies:
  - `go mod download`
- Install `mockgen` (required on PATH):
  - `go install github.com/uber-go/mock/mockgen@latest`

## Development Workflow

- Run the tool (interactive prompts):
  - `go run .`
- Run with explicit inputs:
  - `go run . --search /path/to/search --output /path/to/output`
  - `go run . -S /path/to/search -O /path/to/output`
- Non-interactive selection:
  - `-A` or `--all` to accept all interfaces
  - `-P` or `--skip-path-prompt` to skip per-interface path prompts

## Testing Instructions

- No automated tests are currently present.
- For changes, run `go test ./...` to ensure the code compiles (will be a no-op if no tests exist).

## Code Style

- Use standard Go formatting:
  - `gofmt -w .`
- Keep public APIs stable unless required for a change.
- Favor small, focused functions in `generator/` and `tui/` to keep TUI flow readable.

## Build and Deployment

- Build a local binary:
  - `go build .`
- Optionally move the binary to your Go bin directory:
  - `mv deepsearch-mockgen-cli "$GOPATH/bin"`

## Pull Request Guidelines

- Keep PRs small and focused.
- Ensure `go test ./...` and `gofmt -w .` are clean before submitting.

## Troubleshooting

- If mock generation fails, confirm `mockgen` is installed and on your PATH:
  - `which mockgen`
- If no interfaces are found, verify `--search` points to a Go module or package tree with `.go` files (non-test files).
