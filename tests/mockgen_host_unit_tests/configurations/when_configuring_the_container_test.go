package configurations_test

import (
	"testing"

	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/use_cases"
	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen_cli/controllers"
	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen_cli/flags"
	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen_host/configurations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWhenConfiguringTheContainer(t *testing.T) {
	t.Parallel()

	t.Run("Given the parsed command line", func(t *testing.T) {
		t.Parallel()

		t.Run("Should resolve the use case and its controller", func(t *testing.T) {
			t.Parallel()

			container := configurations.ConfigureDI(
				flags.Parse("deepsearch-mockgen-cli", []string{"-S", "src", "--output", "out", "-A", "--skip-path-prompt"}),
			)

			require.NoError(t, container.Invoke(func(
				generateMocks *use_cases.GenerateMocks,
				controller *controllers.GenerateMocksController,
			) {
				assert.NotNil(t, generateMocks)
				assert.NotNil(t, controller)
			}))
		})

		t.Run("Should accept both short and long flags", func(t *testing.T) {
			t.Parallel()

			short := flags.Parse("deepsearch-mockgen-cli", []string{"-S", "src", "-O", "out", "-A", "-P"})
			long := flags.Parse("deepsearch-mockgen-cli", []string{"--search", "src", "--output", "out", "--all", "--skip-path-prompt"})

			expected := &flags.MockgenFlags{SearchDir: "src", OutputDir: "out", AcceptAll: true, SkipPathPrompt: true}
			assert.Equal(t, expected, short)
			assert.Equal(t, expected, long)
		})
	})
}
