package configurations

import (
	mockgen_configurations "github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/configurations"
	cli_configurations "github.com/FournyP/deepsearch-mockgen-cli/src/mockgen_cli/configurations"
	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen_cli/flags"
	golang_configurations "github.com/FournyP/deepsearch-mockgen-cli/src/mockgen_golang/configurations"
	tui_configurations "github.com/FournyP/deepsearch-mockgen-cli/src/mockgen_tui/configurations"
	uber_configurations "github.com/FournyP/deepsearch-mockgen-cli/src/mockgen_uber/configurations"
	"go.uber.org/dig"
)

func ConfigureDI(parsed *flags.MockgenFlags) *dig.Container {
	container := dig.New()

	AddSettingsConfiguration(container, parsed)
	golang_configurations.AddMockgenGolangConfiguration(container)
	uber_configurations.AddMockgenUberConfiguration(container)
	tui_configurations.AddMockgenTuiConfiguration(container)
	mockgen_configurations.AddMockgenConfiguration(container)
	cli_configurations.AddMockgenCliConfiguration(container)

	return container
}
