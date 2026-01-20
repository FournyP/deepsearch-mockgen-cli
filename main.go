package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/FournyP/deepsearch-mockgen/generator"
	"github.com/FournyP/deepsearch-mockgen/tui"
)

func main() {
	// Define CLI flags
	searchDir := flag.String("search", "", "Directory to search for interfaces")
	outputDir := flag.String("output", "", "Directory to save generated mocks")

	var acceptAll bool
	flag.BoolVar(&acceptAll, "A", false, "Generate mocks for all interfaces without prompting")
	flag.BoolVar(&acceptAll, "all", false, "Generate mocks for all interfaces without prompting")

	var skipPathPrompt bool
	flag.BoolVar(&skipPathPrompt, "S", false, "Skip per-interface mock path prompts and use defaults")
	flag.BoolVar(&skipPathPrompt, "skip-path-prompt", false, "Skip per-interface mock path prompts and use defaults")

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
	for iface, mockPath := range finalPaths {
		err := generator.GenerateMock(iface, interfaces[iface], mockPath)
		if err != nil {
			log.Printf("Failed to generate mock for %s: %v\n", iface, err)
		} else {
			log.Printf("Mock generated for %s at %s\n", iface, mockPath)
		}
	}
}
