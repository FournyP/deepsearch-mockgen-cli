package services_test

import (
	"errors"
	"testing"

	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/models"
	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/services"
	interfaces_mocks "github.com/FournyP/deepsearch-mockgen-cli/tests/mockgen_mocks/services_mocks/interfaces_mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type WhenGeneratingAMockBatchTestingSuite struct {
	sut *services.MockBatchGenerator

	mockMockGenerator *interfaces_mocks.MockMockGeneratorInterface
}

func WhenGeneratingAMockBatchBeforeEach(t *testing.T) *WhenGeneratingAMockBatchTestingSuite {
	mockController := gomock.NewController(t)
	mockMockGenerator := interfaces_mocks.NewMockMockGeneratorInterface(mockController)

	return &WhenGeneratingAMockBatchTestingSuite{
		sut:               services.NewMockBatchGenerator(mockMockGenerator),
		mockMockGenerator: mockMockGenerator,
	}
}

func collectUpdates(updates <-chan models.ProgressUpdate) []models.ProgressUpdate {
	collected := make([]models.ProgressUpdate, 0)
	for update := range updates {
		collected = append(collected, update)
	}
	return collected
}

func TestWhenGeneratingAMockBatch(t *testing.T) {
	t.Parallel()

	t.Run("Given several targets", func(t *testing.T) {
		t.Parallel()

		t.Run("Should send one update per target in order then close", func(t *testing.T) {
			t.Parallel()

			suite := WhenGeneratingAMockBatchBeforeEach(t)
			first := models.MockTarget{InterfaceName: "First", SourcePath: "a.go", MockPath: "a_mock.go"}
			second := models.MockTarget{InterfaceName: "Second", SourcePath: "b.go", MockPath: "b_mock.go"}
			failure := errors.New("mockgen execution failed")
			gomock.InOrder(
				suite.mockMockGenerator.EXPECT().Generate(first).Return(nil),
				suite.mockMockGenerator.EXPECT().Generate(second).Return(failure),
			)
			updates := make(chan models.ProgressUpdate)

			go suite.sut.Generate([]models.MockTarget{first, second}, updates)

			assert.Equal(t, []models.ProgressUpdate{
				{Name: "First"},
				{Name: "Second", Err: failure},
			}, collectUpdates(updates))
		})
	})

	t.Run("Given no targets", func(t *testing.T) {
		t.Parallel()

		t.Run("Should close the updates right away", func(t *testing.T) {
			t.Parallel()

			suite := WhenGeneratingAMockBatchBeforeEach(t)
			updates := make(chan models.ProgressUpdate)

			go suite.sut.Generate(nil, updates)

			assert.Empty(t, collectUpdates(updates))
		})
	})
}
