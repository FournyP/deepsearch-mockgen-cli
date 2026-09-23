package services

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/models"
)

const (
	mockgenBinary      = "mockgen"
	defaultPackageName = "mocks"
)

// MockGenerator runs the uber-go/mock mockgen binary (source mode) for one
// target. The mock package is named after the directory holding the mock file.
type MockGenerator struct{}

func NewMockGenerator() *MockGenerator {
	return &MockGenerator{}
}

func (g *MockGenerator) Generate(target models.MockTarget) error {
	segments := strings.Split(target.MockPath, "/")
	outputDir := strings.Join(segments[:len(segments)-1], "/")

	if err := createDirIfNotExist(outputDir); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	cmd := exec.Command(
		mockgenBinary,
		"-source="+target.SourcePath,
		"-destination="+target.MockPath,
		"-package="+packageName(segments),
	)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("mockgen execution failed: %w", err)
	}
	return nil
}

func packageName(segments []string) string {
	if len(segments) > 1 {
		return segments[len(segments)-2]
	}
	return defaultPackageName
}

func createDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}
	return nil
}
