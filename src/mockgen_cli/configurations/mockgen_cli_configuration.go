package configurations

import (
	mockgen_interfaces "github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/services/interfaces"
	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen_cli/controllers"
	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen_cli/services"
	"go.uber.org/dig"
)

func AddMockgenCliConfiguration(container *dig.Container) {
	provide(
		container,
		services.NewMockgenReporter,
		dig.As(new(mockgen_interfaces.MockgenReporterInterface)),
	)
	provide(container, controllers.NewGenerateMocksController)
}

func provide(container *dig.Container, constructor any, options ...dig.ProvideOption) {
	if err := container.Provide(constructor, options...); err != nil {
		panic(err)
	}
}
