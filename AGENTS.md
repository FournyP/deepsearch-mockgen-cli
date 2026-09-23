# AGENTS.md

## Project Overview

Deepsearch-mockgen-cli is a Go CLI that scans a directory for Go interfaces and generates mocks using the `mockgen` binary (from `github.com/uber-go/mock`). It provides a TUI flow for selecting interfaces and confirming output paths, plus CLI flags for non-interactive runs.

The code follows a hexagonal (ports and adapters) layout wired with `go.uber.org/dig`:

```
src/
  common/utils/            shared helpers (ToSnakeCase)
  mockgen/                 domain core — no I/O, no framework
    models/                MockTarget, ProgressUpdate
    settings/              GenerationSettings (dirs), InteractionSettings (-A, -P)
    services/interfaces/   ports: InterfaceFinder, InterfaceSelector, Prompter,
                           MockPathBuilder, MockGenerator, MockBatchGenerator,
                           ProgressReporter, MockgenReporter
    services/              MockPathBuilder (default mock path), MockBatchGenerator
    use_cases/             GenerateMocks (orchestrates the whole run)
    configurations/        dig registrations of the domain
  mockgen_golang/          adapter: go/ast interface discovery (InterfaceFinder)
  mockgen_uber/            adapter: runs the uber-go/mock `mockgen` binary (MockGenerator)
  mockgen_tui/             adapter: Bubble Tea prompts, selector and progress bar
  mockgen_cli/             adapter: flags, usage, controller, log reporter
  mockgen_host/            composition root: parses flags, builds the dig container, runs
tests/
  *_unit_tests/            when_*_test.go, Given/Should subtests, gomock
  mockgen_mocks/           generated with this very CLI (`make mocks`)
```

Dependency rule: adapters (`mockgen_*`) depend on the domain (`mockgen`), never the
reverse. The domain only talks to the outside world through the interfaces in
`src/mockgen/services/interfaces`. To add a feature, add a port there, implement it in
an adapter, register both in the matching `configurations` package, and consume the
port from a use case.

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

- Unit tests live under `tests/*_unit_tests`: `make test` (or `go test ./...`).
- After changing a port in `src/mockgen/services/interfaces`, regenerate mocks: `make mocks`.

## Code Style

- Use standard Go formatting:
  - `gofmt -w .`
- Keep public APIs stable unless required for a change.
- One struct per responsibility; business flow stays in `src/mockgen/use_cases`, technology in adapters.

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
