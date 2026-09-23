package services_test

import (
	"errors"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"

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

func (s *WhenGeneratingAMockBatchTestingSuite) run(targets []models.MockTarget) []models.ProgressUpdate {
	updates := make(chan models.ProgressUpdate)
	go s.sut.Generate(targets, updates)

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

		t.Run("Should send one update per target then close", func(t *testing.T) {
			t.Parallel()

			suite := WhenGeneratingAMockBatchBeforeEach(t)
			first := models.MockTarget{InterfaceName: "First", SourcePath: "a.go", MockPath: "a_mock.go"}
			second := models.MockTarget{InterfaceName: "Second", SourcePath: "b.go", MockPath: "b_mock.go"}
			failure := errors.New("mockgen execution failed")
			suite.mockMockGenerator.EXPECT().Generate(first).Return(nil)
			suite.mockMockGenerator.EXPECT().Generate(second).Return(failure)

			assert.ElementsMatch(t, []models.ProgressUpdate{
				{Name: "First"},
				{Name: "Second", Err: failure},
			}, suite.run([]models.MockTarget{first, second}))
		})

		t.Run("Should generate distinct mock paths in parallel", func(t *testing.T) {
			t.Parallel()

			if runtime.NumCPU() < 2 {
				t.Skip("parallel generation needs at least 2 CPUs")
			}

			suite := WhenGeneratingAMockBatchBeforeEach(t)
			var started sync.WaitGroup
			started.Add(2)
			suite.mockMockGenerator.EXPECT().
				Generate(gomock.Any()).
				DoAndReturn(func(models.MockTarget) error {
					started.Done()
					if !waitTimeout(&started, 5*time.Second) {
						return errors.New("generations did not overlap")
					}
					return nil
				}).
				Times(2)

			updates := suite.run([]models.MockTarget{
				{InterfaceName: "First", MockPath: "a_mock.go"},
				{InterfaceName: "Second", MockPath: "b_mock.go"},
			})

			for _, update := range updates {
				assert.NoError(t, update.Err)
			}
		})
	})

	t.Run("Given targets sharing a mock path", func(t *testing.T) {
		t.Parallel()

		t.Run("Should generate them one after another in order", func(t *testing.T) {
			t.Parallel()

			suite := WhenGeneratingAMockBatchBeforeEach(t)
			first := models.MockTarget{InterfaceName: "First", MockPath: "out/shared_mock.go"}
			second := models.MockTarget{InterfaceName: "Second", MockPath: "./out/shared_mock.go"}
			var running atomic.Int32
			overlapped := atomic.Bool{}
			generate := func(models.MockTarget) error {
				if running.Add(1) > 1 {
					overlapped.Store(true)
				}
				time.Sleep(20 * time.Millisecond)
				running.Add(-1)
				return nil
			}
			gomock.InOrder(
				suite.mockMockGenerator.EXPECT().Generate(first).DoAndReturn(generate),
				suite.mockMockGenerator.EXPECT().Generate(second).DoAndReturn(generate),
			)

			assert.Equal(t, []models.ProgressUpdate{
				{Name: "First"},
				{Name: "Second"},
			}, suite.run([]models.MockTarget{first, second}))
			assert.False(t, overlapped.Load())
		})
	})

	t.Run("Given no targets", func(t *testing.T) {
		t.Parallel()

		t.Run("Should close the updates right away", func(t *testing.T) {
			t.Parallel()

			suite := WhenGeneratingAMockBatchBeforeEach(t)

			assert.Empty(t, suite.run(nil))
		})
	})
}

func waitTimeout(group *sync.WaitGroup, timeout time.Duration) bool {
	done := make(chan struct{})
	go func() {
		group.Wait()
		close(done)
	}()

	select {
	case <-done:
		return true
	case <-time.After(timeout):
		return false
	}
}
