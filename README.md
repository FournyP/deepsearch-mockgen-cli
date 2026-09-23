# Deepsearch-mockgen-cli

This is a simple tool to generate mocks for Go interfaces. It uses the `mockgen` tool from the `[github.com/golang/mock](https://github.com/uber-go/mock)` package.

## Installation

### Option 1: `go install`

```bash
go install github.com/FournyP/deepsearch-mockgen-cli@latest
```

This installs the latest tagged release into `$GOPATH/bin` (or `$HOME/go/bin`). Make sure that directory is on your `PATH`.

### Option 2: Prebuilt binaries

Prebuilt binaries for Linux, macOS, and Windows (both `amd64` and `arm64`) are attached to every [GitHub release](https://github.com/FournyP/deepsearch-mockgen-cli/releases).

Download the archive matching your platform, extract it, and move the `deepsearch-mockgen-cli` binary somewhere on your `PATH`:

```bash
# Example for Linux amd64
curl -L -o deepsearch-mockgen-cli.tar.gz \
  https://github.com/FournyP/deepsearch-mockgen-cli/releases/latest/download/deepsearch-mockgen-cli_<version>_linux_x86_64.tar.gz
tar -xzf deepsearch-mockgen-cli.tar.gz
mv deepsearch-mockgen-cli /usr/local/bin/
```

On Windows, download the `.zip` archive instead and extract the `.exe`.

### Option 3: Build from source

```bash
go build .
mv deepsearch-mockgen-cli $GOPATH/bin
```

## Usage

```bash
deepsearch-mockgen-cli
```

Options:

```
-S, --search <dir>       Directory to search for interfaces
-O, --output <dir>       Directory to save generated mocks
-A, --all                Generate mocks for all interfaces without prompting
-P, --skip-path-prompt   Skip per-interface mock path prompts and use defaults
```

Missing directories are prompted for. Mocks mirror the source tree under the output
directory, each directory suffixed with `_mocks` (e.g. `src/pkg/saver.go` with
`-S src -O tests` gives `tests/pkg_mocks/saver_mock.go`).

## Development

```bash
make test   # unit tests
make mocks  # regenerate test mocks with this CLI
```

Architecture (hexagonal, wired with `go.uber.org/dig`) is described in [AGENTS.md](AGENTS.md).
