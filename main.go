package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/FournyP/deepsearch-mockgen/generator"
	"github.com/FournyP/deepsearch-mockgen/tui"
)

func main() {
	// Define CLI flags
	searchDir := flag.String("search", "", "Directory to search for interfaces")
	outputDir := flag.String("output", "", "Directory to save generated mocks")
	flag.StringVar(searchDir, "S", "", "Directory to search for interfaces")
	flag.StringVar(outputDir, "O", "", "Directory to save generated mocks")

	var acceptAll bool
	flag.BoolVar(&acceptAll, "A", false, "Generate mocks for all interfaces without prompting")
	flag.BoolVar(&acceptAll, "all", false, "Generate mocks for all interfaces without prompting")

	var skipPathPrompt bool
	flag.BoolVar(&skipPathPrompt, "P", false, "Skip per-interface mock path prompts and use defaults")
	flag.BoolVar(&skipPathPrompt, "skip-path-prompt", false, "Skip per-interface mock path prompts and use defaults")

	flag.Usage = func() {
		out := flag.CommandLine.Output()
		fmt.Fprintf(out, "Usage: %s [options]\n\n", os.Args[0])
		fmt.Fprintln(out, "Deepsearch-mockgen deeply scans a directory tree for Go interfaces and generates mocks using mockgen.")
		fmt.Fprintln(out, "Mocks are written under the output directory following the same relative tree as the interface source files.")
		fmt.Fprintln(out, "Options:")
		fmt.Fprintln(out, "  -S, --search <dir>       Directory to search for interfaces")
		fmt.Fprintln(out, "  -O, --output <dir>       Directory to save generated mocks")
		fmt.Fprintln(out, "  -A, --all                Generate mocks for all interfaces without prompting")
		fmt.Fprintln(out, "  -P, --skip-path-prompt   Skip per-interface mock path prompts and use defaults")
		fmt.Fprintln(out, "  -h, --help               Show help")
	}

	// Parse flags
	flag.Parse()

	// Prompt for missing values (uses TUI prompt helpers with stdin fallback)
	if *searchDir == "" {
		var err error
		*searchDir, err = tui.RunTextInputPrompt("Enter the search directory:", "")
		if err != nil {
			log.Fatal(err)
		}
	}

	if *outputDir == "" {
		var err error
		*outputDir, err = tui.RunTextInputPrompt("Enter the output directory:", "")
		if err != nil {
			log.Fatal(err)
		}
	}

	// Find interfaces
	interfaces, err := generator.FindInterfaces(*searchDir)
	if err != nil {
		log.Fatal(err)
	}

	if len(interfaces) == 0 {
		log.Println("No interfaces found")
		return
	}

	// Let user select interfaces via TUI list (or accept all when `-A/--all` provided)
	var selected map[string]string
	if acceptAll {
		selected = make(map[string]string)
		for iface := range interfaces {
			selected[iface] = *outputDir
		}
	} else {
		sel, err := tui.RunInterfaceSelector(interfaces, *outputDir)
		if err != nil {
			log.Fatal(err)
		}
		selected = sel
	}

	// For each selected interface, compute default path and confirm/modify
	finalPaths := make(map[string]string)
	for iface := range selected {
		ifacePath := interfaces[iface]
		mockPath := generator.ComputeMockPath(*searchDir, *outputDir, ifacePath, iface)

		if !skipPathPrompt {
			var err error
			mockPath, err = tui.RunTextInputPrompt(
				fmt.Sprintf("Mock path for %s:", iface),
				mockPath,
			)
			if err != nil {
				log.Fatal(err)
			}
		}

		finalPaths[iface] = mockPath
	}

	// Generate mocks
	// Generate mocks and show progress via TUI
	updates := make(chan tui.ProgressUpdate)
	go func() {
		for iface, mockPath := range finalPaths {
			err := generator.GenerateMock(iface, interfaces[iface], mockPath)
			updates <- tui.ProgressUpdate{
				Name: iface,
				Err:  err,
			}
		}
		close(updates)
	}()

	if err := tui.RunProgress(len(finalPaths), updates); err != nil {
		log.Printf("Progress UI error: %v", err)
	}
}
