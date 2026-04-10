# GitHub Copilot Instructions

## Priority Guidelines

When generating code for this repository:

1. **Version Compatibility**: Use Go 1.24 features only (see go.mod).
2. **Context Files**: There are no existing files in .github/copilot; rely on the codebase.
3. **Codebase Patterns**: Follow established patterns in `main`, `generator`, and `tui` packages.
4. **Architectural Consistency**: Keep the single-module CLI structure with small packages (`generator/`, `tui/`).
5. **Code Quality**: Favor maintainability and clarity consistent with existing code.

## Technology Version Detection

- **Language**: Go 1.24.0 (go.mod).
- **Primary libraries**:
  - `github.com/charmbracelet/bubbletea` v1.3.10
  - `github.com/charmbracelet/bubbles` v0.21.0
- **External tool**: `mockgen` from `github.com/uber-go/mock` (installed via `go install`).

Generate code compatible with these versions only.

## Codebase Structure and Architecture

- **Entry point**: `main.go` handles CLI flags, prompts, and orchestration.
- **Generator layer**: `generator/` provides interface discovery, mock path calculation, and `mockgen` invocation.
- **TUI layer**: `tui/` hosts Bubble Tea models and prompt runners.

Keep new logic in the appropriate package and avoid adding new layers unless required.

## Codebase Patterns to Follow

### Naming and Organization

- Package names are short and lowercase (`generator`, `tui`).
- Functions are small and focused (e.g., `FindInterfaces`, `ComputeMockPath`, `RunProgress`).
- Use simple data structs for Bubble Tea models (`textInputModel`, `progressModel`).

### Error Handling

- Return errors up the call chain and wrap with `fmt.Errorf` where appropriate (e.g., `CreateDirIfNotExist`, `GenerateMock`).
- In `main.go`, errors are handled with `log.Fatal` or `log.Printf`.

### CLI and TUI Behavior

- Flags use `flag` package with short and long variants.
- TUI flows use Bubble Tea programs and return a model or error.
- Cancellation uses `ctrl+c` handling in TUI models with a `Cancelled` boolean.

## Documentation Requirements

- Comments are minimal and focused on exported functions and key behavior.
- Match existing comment style (single-line `//` comments above functions).

## Testing Approach

- No tests exist today. If adding tests, keep them minimal and aligned with standard Go `testing` conventions.

## General Best Practices (Codebase-Consistent)

- Keep functions focused and avoid complex abstractions.
- Prefer standard library packages where possible.
- Maintain existing CLI/TUI flow structure and UI text conventions.
