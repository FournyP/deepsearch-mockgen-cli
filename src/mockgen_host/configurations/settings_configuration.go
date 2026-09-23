package configurations

import (
	mockgen_settings "github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/settings"
	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen_cli/flags"
	"go.uber.org/dig"
)

func AddSettingsConfiguration(container *dig.Container, parsed *flags.MockgenFlags) {
	provide(container, func() *mockgen_settings.GenerationSettings {
		return mockgen_settings.NewGenerationSettings(parsed.SearchDir, parsed.OutputDir)
	})
	provide(container, func() *mockgen_settings.InteractionSettings {
		return mockgen_settings.NewInteractionSettings(parsed.AcceptAll, parsed.SkipPathPrompt)
	})
}

func provide(container *dig.Container, constructor any, options ...dig.ProvideOption) {
	if err := container.Provide(constructor, options...); err != nil {
		panic(err)
	}
}
