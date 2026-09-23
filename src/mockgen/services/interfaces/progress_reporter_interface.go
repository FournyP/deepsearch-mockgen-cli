package interfaces

import "github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/models"

type ProgressReporterInterface interface {
	Report(total int, updates <-chan models.ProgressUpdate) error
}
