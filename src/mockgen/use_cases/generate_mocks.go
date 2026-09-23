package use_cases

import (
	"fmt"
	"sort"

	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/models"
	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/services/interfaces"
	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/settings"
)

const (
	searchDirQuestion = "Enter the search directory:"
	outputDirQuestion = "Enter the output directory:"
)

type GenerateMocks struct {
	interfaceFinder     interfaces.InterfaceFinderInterface
	interfaceSelector   interfaces.InterfaceSelectorInterface
	prompter            interfaces.PrompterInterface
	mockPathBuilder     interfaces.MockPathBuilderInterface
	mockBatchGenerator  interfaces.MockBatchGeneratorInterface
	progressReporter    interfaces.ProgressReporterInterface
	mockgenReporter     interfaces.MockgenReporterInterface
	generationSettings  *settings.GenerationSettings
	interactionSettings *settings.InteractionSettings
}

func NewGenerateMocks(
	interfaceFinder interfaces.InterfaceFinderInterface,
	interfaceSelector interfaces.InterfaceSelectorInterface,
	prompter interfaces.PrompterInterface,
	mockPathBuilder interfaces.MockPathBuilderInterface,
	mockBatchGenerator interfaces.MockBatchGeneratorInterface,
	progressReporter interfaces.ProgressReporterInterface,
	mockgenReporter interfaces.MockgenReporterInterface,
	generationSettings *settings.GenerationSettings,
	interactionSettings *settings.InteractionSettings,
) *GenerateMocks {
	return &GenerateMocks{
		interfaceFinder:     interfaceFinder,
		interfaceSelector:   interfaceSelector,
		prompter:            prompter,
		mockPathBuilder:     mockPathBuilder,
		mockBatchGenerator:  mockBatchGenerator,
		progressReporter:    progressReporter,
		mockgenReporter:     mockgenReporter,
		generationSettings:  generationSettings,
		interactionSettings: interactionSettings,
	}
}

func (u *GenerateMocks) Execute() error {
	searchDir, err := u.resolveDir(u.generationSettings.SearchDir, searchDirQuestion)
	if err != nil {
		return err
	}

	outputDir, err := u.resolveDir(u.generationSettings.OutputDir, outputDirQuestion)
	if err != nil {
		return err
	}

	found, err := u.interfaceFinder.Find(searchDir)
	if err != nil {
		return err
	}

	if len(found) == 0 {
		u.mockgenReporter.NoInterfacesFound()
		return nil
	}

	selected, err := u.selectInterfaces(found)
	if err != nil {
		return err
	}

	targets, err := u.buildTargets(searchDir, outputDir, found, selected)
	if err != nil {
		return err
	}

	u.generate(targets)
	return nil
}

// resolveDir prompts for a directory when none was given on the command line.
func (u *GenerateMocks) resolveDir(dir, question string) (string, error) {
	if dir != "" {
		return dir, nil
	}
	return u.prompter.Prompt(question, "")
}

func (u *GenerateMocks) selectInterfaces(found map[string]string) ([]string, error) {
	if u.interactionSettings.AcceptAll {
		return sortedNames(found), nil
	}

	selected, err := u.interfaceSelector.Select(found)
	if err != nil {
		return nil, err
	}
	sort.Strings(selected)
	return selected, nil
}

// buildTargets computes the default mock path of each selected interface and,
// unless path prompts are skipped, lets the user confirm or modify it.
func (u *GenerateMocks) buildTargets(
	searchDir, outputDir string,
	found map[string]string,
	selected []string,
) ([]models.MockTarget, error) {
	targets := make([]models.MockTarget, 0, len(selected))
	for _, name := range selected {
		sourcePath := found[name]
		mockPath := u.mockPathBuilder.Build(searchDir, outputDir, sourcePath, name)

		if !u.interactionSettings.SkipPathPrompt {
			var err error
			mockPath, err = u.prompter.Prompt(fmt.Sprintf("Mock path for %s:", name), mockPath)
			if err != nil {
				return nil, err
			}
		}

		targets = append(targets, models.MockTarget{
			InterfaceName: name,
			SourcePath:    sourcePath,
			MockPath:      mockPath,
		})
	}
	return targets, nil
}

// generate runs the generation in the background while the progress reporter
// consumes its updates. A progress UI failure is reported, not returned.
func (u *GenerateMocks) generate(targets []models.MockTarget) {
	updates := make(chan models.ProgressUpdate)
	go u.mockBatchGenerator.Generate(targets, updates)

	if err := u.progressReporter.Report(len(targets), updates); err != nil {
		u.mockgenReporter.ProgressFailed(err)
	}
}

func sortedNames(found map[string]string) []string {
	names := make([]string, 0, len(found))
	for name := range found {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}
