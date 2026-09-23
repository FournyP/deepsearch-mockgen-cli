package configurations

import (
	mockgen_interfaces "github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/services/interfaces"
	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen_uber/services"
	"go.uber.org/dig"
)

func AddMockgenUberConfiguration(container *dig.Container) {
	provide(
		container,
		services.NewMockGenerator,
		dig.As(new(mockgen_interfaces.MockGeneratorInterface)),
	)
}

func provide(container *dig.Container, constructor any, options ...dig.ProvideOption) {
	if err := container.Provide(constructor, options...); err != nil {
		panic(err)
	}
}
