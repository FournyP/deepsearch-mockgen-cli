package interfaces

import "github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/models"

type MockBatchGeneratorInterface interface {
	// Generate generates every target, sends one update per target, then closes updates.
	Generate(targets []models.MockTarget, updates chan<- models.ProgressUpdate)
}
