package services

import (
	"path/filepath"
	"runtime"
	"sync"

	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/models"
	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/services/interfaces"
)

// MockBatchGenerator generates targets in parallel, up to one generation per
// CPU. Targets sharing a mock path are generated one after another, in their
// original order, so they never write the same file concurrently.
type MockBatchGenerator struct {
	mockGenerator interfaces.MockGeneratorInterface
}

func NewMockBatchGenerator(mockGenerator interfaces.MockGeneratorInterface) *MockBatchGenerator {
	return &MockBatchGenerator{mockGenerator: mockGenerator}
}

func (g *MockBatchGenerator) Generate(
	targets []models.MockTarget,
	updates chan<- models.ProgressUpdate,
) {
	defer close(updates)

	groups := groupByMockPath(targets)
	jobs := make(chan []models.MockTarget)

	var workers sync.WaitGroup
	for range min(runtime.NumCPU(), len(groups)) {
		workers.Add(1)
		go func() {
			defer workers.Done()
			for group := range jobs {
				g.generateGroup(group, updates)
			}
		}()
	}

	for _, group := range groups {
		jobs <- group
	}
	close(jobs)
	workers.Wait()
}

func (g *MockBatchGenerator) generateGroup(
	group []models.MockTarget,
	updates chan<- models.ProgressUpdate,
) {
	for _, target := range group {
		updates <- models.ProgressUpdate{
			Name: target.InterfaceName,
			Err:  g.mockGenerator.Generate(target),
		}
	}
}

// groupByMockPath groups targets writing the same file, keeping first-seen order.
func groupByMockPath(targets []models.MockTarget) [][]models.MockTarget {
	indexes := make(map[string]int)
	groups := make([][]models.MockTarget, 0, len(targets))
	for _, target := range targets {
		key := filepath.Clean(target.MockPath)
		index, ok := indexes[key]
		if !ok {
			index = len(groups)
			indexes[key] = index
			groups = append(groups, nil)
		}
		groups[index] = append(groups[index], target)
	}
	return groups
}
