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
