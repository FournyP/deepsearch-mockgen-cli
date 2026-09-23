package interfaces

import "github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/models"

type MockBatchGeneratorInterface interface {
	// Generate generates every target in parallel, sends one update per target
	// (in completion order), then closes updates.
	Generate(targets []models.MockTarget, updates chan<- models.ProgressUpdate)
}
