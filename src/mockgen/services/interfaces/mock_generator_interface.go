package interfaces

import "github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/models"

type MockGeneratorInterface interface {
	Generate(target models.MockTarget) error
}
