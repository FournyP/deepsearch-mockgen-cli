package flags

import (
	"flag"
	"os"
)

type MockgenFlags struct {
	SearchDir      string
	OutputDir      string
	AcceptAll      bool
	SkipPathPrompt bool
}

func Parse(binaryName string, args []string) *MockgenFlags {
	parsed := &MockgenFlags{}

	set := flag.NewFlagSet(binaryName, flag.ExitOnError)
	set.Usage = func() { Usage(set.Output(), binaryName) }

	set.StringVar(&parsed.SearchDir, "search", "", "Directory to search for interfaces")
	set.StringVar(&parsed.SearchDir, "S", "", "Directory to search for interfaces")
	set.StringVar(&parsed.OutputDir, "output", "", "Directory to save generated mocks")
	set.StringVar(&parsed.OutputDir, "O", "", "Directory to save generated mocks")
	set.BoolVar(&parsed.AcceptAll, "all", false, "Generate mocks for all interfaces without prompting")
	set.BoolVar(&parsed.AcceptAll, "A", false, "Generate mocks for all interfaces without prompting")
	set.BoolVar(&parsed.SkipPathPrompt, "skip-path-prompt", false, "Skip per-interface mock path prompts and use defaults")
	set.BoolVar(&parsed.SkipPathPrompt, "P", false, "Skip per-interface mock path prompts and use defaults")

	if err := set.Parse(args); err != nil {
		os.Exit(2)
	}
	return parsed
}
