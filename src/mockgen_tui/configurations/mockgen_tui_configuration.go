package configurations

import (
	mockgen_interfaces "github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/services/interfaces"
	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen_tui/services"
	"go.uber.org/dig"
)

func AddMockgenTuiConfiguration(container *dig.Container) {
	provide(
		container,
		services.NewInterfaceSelector,
		dig.As(new(mockgen_interfaces.InterfaceSelectorInterface)),
	)
	provide(container, services.NewPrompter, dig.As(new(mockgen_interfaces.PrompterInterface)))
	provide(
		container,
		services.NewProgressReporter,
		dig.As(new(mockgen_interfaces.ProgressReporterInterface)),
	)
}

func provide(container *dig.Container, constructor any, options ...dig.ProvideOption) {
	if err := container.Provide(constructor, options...); err != nil {
		panic(err)
	}
}
