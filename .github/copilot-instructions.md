# GitHub Copilot Instructions

## Priority Guidelines

When generating code for this repository:

1. **Version Compatibility**: Use Go 1.24 features only (see go.mod).
2. **Context Files**: There are no existing files in .github/copilot; rely on the codebase.
3. **Codebase Patterns**: Follow established patterns in the `src/` domain and adapter packages.
4. **Architectural Consistency**: Keep the hexagonal layout: domain in `src/mockgen`, adapters in `src/mockgen_*`, wiring in `configurations` packages.
5. **Code Quality**: Favor maintainability and clarity consistent with existing code.

## Technology Version Detection

- **Language**: Go 1.24.0 (go.mod).
- **Primary libraries**:
  - `github.com/charmbracelet/bubbletea` v1.3.10
  - `github.com/charmbracelet/bubbles` v0.21.0
  - `go.uber.org/dig` v1.19.0 (dependency injection)
  - `go.uber.org/mock` v0.6.0 and `github.com/stretchr/testify` (tests)
- **External tool**: `mockgen` from `github.com/uber-go/mock` (installed via `go install`).

Generate code compatible with these versions only.

## Codebase Structure and Architecture

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

- `main.go` only calls `mockgen_host.Run()`.
- Domain code never imports an adapter; it depends on the ports in `src/mockgen/services/interfaces`.
- Each package with registrations exposes an `Add<Name>Configuration(container *dig.Container)` function.

## Codebase Patterns to Follow

### Naming and Organization

- Ports are named `<Thing>Interface` in `<thing>_interface.go`; implementations are `<Thing>` built by `New<Thing>`.
- Use cases expose `Execute()`; controllers only translate CLI input to use case calls.
- Bubble Tea models live in `src/mockgen_tui/models` (`TextInputModel`, `InterfaceSelectorModel`, `ProgressModel`).

### Error Handling

- Return errors up the call chain and wrap with `fmt.Errorf` where appropriate (e.g., `MockGenerator.Generate`).
- `mockgen_host.Run` handles the final error with `log.Fatal`; non-fatal messages go through `MockgenReporterInterface`.

### CLI and TUI Behavior

- Flags use `flag` package with short and long variants.
- TUI flows use Bubble Tea programs and return a model or error.
- Cancellation uses `ctrl+c` handling in TUI models with a `Cancelled` boolean.

## Documentation Requirements

- Comments are minimal and focused on exported functions and key behavior.
- Match existing comment style (single-line `//` comments above functions).

## Testing Approach

- Tests live in `tests/*_unit_tests/**/when_*_test.go` with `Given ...`/`Should ...` subtests, a `...BeforeEach` suite builder, testify and gomock.
- Mocks are generated into `tests/mockgen_mocks` with `make mocks`.

## General Best Practices (Codebase-Consistent)

- Keep functions focused and avoid complex abstractions.
- Prefer standard library packages where possible.
- Maintain existing CLI/TUI flow structure and UI text conventions.
