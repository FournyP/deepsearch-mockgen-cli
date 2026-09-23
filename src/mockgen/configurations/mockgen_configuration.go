package configurations

import (
	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/services"
	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/services/interfaces"
	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/use_cases"
	"go.uber.org/dig"
)

func AddMockgenConfiguration(container *dig.Container) {
	provide(container, services.NewMockPathBuilder, dig.As(new(interfaces.MockPathBuilderInterface)))
	provide(container, services.NewMockBatchGenerator, dig.As(new(interfaces.MockBatchGeneratorInterface)))

	provide(container, use_cases.NewGenerateMocks)
}

func provide(container *dig.Container, constructor any, options ...dig.ProvideOption) {
	if err := container.Provide(constructor, options...); err != nil {
		panic(err)
	}
}
