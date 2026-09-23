package flags

import (
	"fmt"
	"io"
)

func Usage(out io.Writer, binaryName string) {
	fmt.Fprintf(out, "Usage: %s [options]\n\n", binaryName)
	fmt.Fprintln(out, "Deepsearch-mockgen-cli deeply scans a directory tree for Go interfaces and generates mocks using mockgen.")
	fmt.Fprintln(out, "Mocks are written under the output directory following the same relative tree as the interface source files.")
	fmt.Fprintln(out, "Options:")
	fmt.Fprintln(out, "  -S, --search <dir>       Directory to search for interfaces")
	fmt.Fprintln(out, "  -O, --output <dir>       Directory to save generated mocks")
	fmt.Fprintln(out, "  -A, --all                Generate mocks for all interfaces without prompting")
	fmt.Fprintln(out, "  -P, --skip-path-prompt   Skip per-interface mock path prompts and use defaults")
	fmt.Fprintln(out, "  -h, --help               Show help")
}
