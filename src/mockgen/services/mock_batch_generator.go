package services

import (
	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/models"
	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/services/interfaces"
)

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

	for _, target := range targets {
		updates <- models.ProgressUpdate{
			Name: target.InterfaceName,
			Err:  g.mockGenerator.Generate(target),
		}
	}
}
