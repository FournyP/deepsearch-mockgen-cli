package configurations

import (
	mockgen_interfaces "github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/services/interfaces"
	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen_golang/services"
	"go.uber.org/dig"
)

func AddMockgenGolangConfiguration(container *dig.Container) {
	provide(
		container,
		services.NewInterfaceFinder,
		dig.As(new(mockgen_interfaces.InterfaceFinderInterface)),
	)
}

func provide(container *dig.Container, constructor any, options ...dig.ProvideOption) {
	if err := container.Provide(constructor, options...); err != nil {
		panic(err)
	}
}
