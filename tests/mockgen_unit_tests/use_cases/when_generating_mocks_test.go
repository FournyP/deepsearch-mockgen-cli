package use_cases_test

import (
	"errors"
	"testing"

	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/models"
	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/settings"
	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/use_cases"
	interfaces_mocks "github.com/FournyP/deepsearch-mockgen-cli/tests/mockgen_mocks/services_mocks/interfaces_mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

type WhenGeneratingMocksTestingSuite struct {
	sut *use_cases.GenerateMocks

	mockInterfaceFinder    *interfaces_mocks.MockInterfaceFinderInterface
	mockInterfaceSelector  *interfaces_mocks.MockInterfaceSelectorInterface
	mockPrompter           *interfaces_mocks.MockPrompterInterface
	mockMockPathBuilder    *interfaces_mocks.MockMockPathBuilderInterface
	mockMockBatchGenerator *interfaces_mocks.MockMockBatchGeneratorInterface
	mockProgressReporter   *interfaces_mocks.MockProgressReporterInterface
	mockMockgenReporter    *interfaces_mocks.MockMockgenReporterInterface
}

func WhenGeneratingMocksBeforeEach(
	t *testing.T,
	generationSettings *settings.GenerationSettings,
	interactionSettings *settings.InteractionSettings,
) *WhenGeneratingMocksTestingSuite {
	mockController := gomock.NewController(t)
	suite := &WhenGeneratingMocksTestingSuite{
		mockInterfaceFinder:    interfaces_mocks.NewMockInterfaceFinderInterface(mockController),
		mockInterfaceSelector:  interfaces_mocks.NewMockInterfaceSelectorInterface(mockController),
		mockPrompter:           interfaces_mocks.NewMockPrompterInterface(mockController),
		mockMockPathBuilder:    interfaces_mocks.NewMockMockPathBuilderInterface(mockController),
		mockMockBatchGenerator: interfaces_mocks.NewMockMockBatchGeneratorInterface(mockController),
		mockProgressReporter:   interfaces_mocks.NewMockProgressReporterInterface(mockController),
		mockMockgenReporter:    interfaces_mocks.NewMockMockgenReporterInterface(mockController),
	}
	suite.sut = use_cases.NewGenerateMocks(
		suite.mockInterfaceFinder,
		suite.mockInterfaceSelector,
		suite.mockPrompter,
		suite.mockMockPathBuilder,
		suite.mockMockBatchGenerator,
		suite.mockProgressReporter,
		suite.mockMockgenReporter,
		generationSettings,
		interactionSettings,
	)
	return suite
}

// expectGeneration captures the targets handed to the batch generator and
// drains its updates through the progress reporter.
func (s *WhenGeneratingMocksTestingSuite) expectGeneration(progressErr error) *[]models.MockTarget {
	generated := new([]models.MockTarget)
	s.mockMockBatchGenerator.EXPECT().
		Generate(gomock.Any(), gomock.Any()).
		DoAndReturn(func(targets []models.MockTarget, updates chan<- models.ProgressUpdate) {
			*generated = targets
			close(updates)
		})
	s.mockProgressReporter.EXPECT().
		Report(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ int, updates <-chan models.ProgressUpdate) error {
			for range updates {
			}
			return progressErr
		})
	return generated
}

func (s *WhenGeneratingMocksTestingSuite) expectDefaultPaths() {
	s.mockMockPathBuilder.EXPECT().
		Build("src", "out", gomock.Any(), gomock.Any()).
		DoAndReturn(func(_, _, _, name string) string { return "out/" + name + "_mock.go" }).
		AnyTimes()
}

var foundInterfaces = map[string]string{
	"Beta":  "src/b.go",
	"Alpha": "src/a.go",
}

func TestWhenGeneratingMocks(t *testing.T) {
	t.Parallel()

	t.Run("Given no directories on the command line", func(t *testing.T) {
		t.Parallel()

		t.Run("Should prompt for the search and output directories", func(t *testing.T) {
			t.Parallel()

			suite := WhenGeneratingMocksBeforeEach(
				t,
				settings.NewGenerationSettings("", ""),
				settings.NewInteractionSettings(true, true),
			)
			gomock.InOrder(
				suite.mockPrompter.EXPECT().Prompt("Enter the search directory:", "").Return("src", nil),
				suite.mockPrompter.EXPECT().Prompt("Enter the output directory:", "").Return("out", nil),
			)
			suite.mockInterfaceFinder.EXPECT().Find("src").Return(map[string]string{}, nil)
			suite.mockMockgenReporter.EXPECT().NoInterfacesFound()

			require.NoError(t, suite.sut.Execute())
		})

		t.Run("Should stop when a prompt is cancelled", func(t *testing.T) {
			t.Parallel()

			suite := WhenGeneratingMocksBeforeEach(
				t,
				settings.NewGenerationSettings("", "out"),
				settings.NewInteractionSettings(true, true),
			)
			suite.mockPrompter.EXPECT().Prompt(gomock.Any(), gomock.Any()).Return("", errors.New("cancelled"))

			assert.ErrorContains(t, suite.sut.Execute(), "cancelled")
		})
	})

	t.Run("Given a search directory without interfaces", func(t *testing.T) {
		t.Parallel()

		t.Run("Should report it and generate nothing", func(t *testing.T) {
			t.Parallel()

			suite := WhenGeneratingMocksBeforeEach(
				t,
				settings.NewGenerationSettings("src", "out"),
				settings.NewInteractionSettings(false, false),
			)
			suite.mockInterfaceFinder.EXPECT().Find("src").Return(map[string]string{}, nil)
			suite.mockMockgenReporter.EXPECT().NoInterfacesFound()

			require.NoError(t, suite.sut.Execute())
		})
	})

	t.Run("Given an unreadable search directory", func(t *testing.T) {
		t.Parallel()

		t.Run("Should propagate the failure", func(t *testing.T) {
			t.Parallel()

			suite := WhenGeneratingMocksBeforeEach(
				t,
				settings.NewGenerationSettings("src", "out"),
				settings.NewInteractionSettings(false, false),
			)
			suite.mockInterfaceFinder.EXPECT().Find("src").Return(nil, errors.New("broken"))

			assert.ErrorContains(t, suite.sut.Execute(), "broken")
		})
	})

	t.Run("Given accept all and skipped path prompts", func(t *testing.T) {
		t.Parallel()

		t.Run("Should generate every interface at its default path in name order", func(t *testing.T) {
			t.Parallel()

			suite := WhenGeneratingMocksBeforeEach(
				t,
				settings.NewGenerationSettings("src", "out"),
				settings.NewInteractionSettings(true, true),
			)
			suite.mockInterfaceFinder.EXPECT().Find("src").Return(foundInterfaces, nil)
			suite.expectDefaultPaths()
			generated := suite.expectGeneration(nil)

			require.NoError(t, suite.sut.Execute())
			assert.Equal(t, []models.MockTarget{
				{InterfaceName: "Alpha", SourcePath: "src/a.go", MockPath: "out/Alpha_mock.go"},
				{InterfaceName: "Beta", SourcePath: "src/b.go", MockPath: "out/Beta_mock.go"},
			}, *generated)
		})
	})

	t.Run("Given an interactive selection", func(t *testing.T) {
		t.Parallel()

		t.Run("Should only generate the selected interfaces at the confirmed paths", func(t *testing.T) {
			t.Parallel()

			suite := WhenGeneratingMocksBeforeEach(
				t,
				settings.NewGenerationSettings("src", "out"),
				settings.NewInteractionSettings(false, false),
			)
			suite.mockInterfaceFinder.EXPECT().Find("src").Return(foundInterfaces, nil)
			suite.mockInterfaceSelector.EXPECT().Select(foundInterfaces).Return([]string{"Beta"}, nil)
			suite.expectDefaultPaths()
			suite.mockPrompter.EXPECT().
				Prompt("Mock path for Beta:", "out/Beta_mock.go").
				Return("custom/beta_mock.go", nil)
			generated := suite.expectGeneration(nil)

			require.NoError(t, suite.sut.Execute())
			assert.Equal(t, []models.MockTarget{
				{InterfaceName: "Beta", SourcePath: "src/b.go", MockPath: "custom/beta_mock.go"},
			}, *generated)
		})

		t.Run("Should stop when the selection is cancelled", func(t *testing.T) {
			t.Parallel()

			suite := WhenGeneratingMocksBeforeEach(
				t,
				settings.NewGenerationSettings("src", "out"),
				settings.NewInteractionSettings(false, false),
			)
			suite.mockInterfaceFinder.EXPECT().Find("src").Return(foundInterfaces, nil)
			suite.mockInterfaceSelector.EXPECT().Select(foundInterfaces).Return(nil, errors.New("cancelled"))

			assert.ErrorContains(t, suite.sut.Execute(), "cancelled")
		})

		t.Run("Should stop when a path prompt is cancelled", func(t *testing.T) {
			t.Parallel()

			suite := WhenGeneratingMocksBeforeEach(
				t,
				settings.NewGenerationSettings("src", "out"),
				settings.NewInteractionSettings(true, false),
			)
			suite.mockInterfaceFinder.EXPECT().Find("src").Return(foundInterfaces, nil)
			suite.expectDefaultPaths()
			suite.mockPrompter.EXPECT().Prompt(gomock.Any(), gomock.Any()).Return("", errors.New("cancelled"))

			assert.ErrorContains(t, suite.sut.Execute(), "cancelled")
		})
	})

	t.Run("Given a failing progress UI", func(t *testing.T) {
		t.Parallel()

		t.Run("Should report the failure without failing the run", func(t *testing.T) {
			t.Parallel()

			suite := WhenGeneratingMocksBeforeEach(
				t,
				settings.NewGenerationSettings("src", "out"),
				settings.NewInteractionSettings(true, true),
			)
			suite.mockInterfaceFinder.EXPECT().Find("src").Return(foundInterfaces, nil)
			suite.expectDefaultPaths()
			progressErr := errors.New("no tty")
			suite.expectGeneration(progressErr)
			suite.mockMockgenReporter.EXPECT().ProgressFailed(progressErr)

			require.NoError(t, suite.sut.Execute())
		})
	})
}
