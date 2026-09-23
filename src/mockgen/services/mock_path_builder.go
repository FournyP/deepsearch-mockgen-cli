package services

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/FournyP/deepsearch-mockgen-cli/src/common/utils"
)

// MockPathBuilder computes the default output path of a mock: the interface
// source tree relative to searchDir is mirrored under outputDir with every
// directory suffixed by "_mocks".
type MockPathBuilder struct{}

func NewMockPathBuilder() *MockPathBuilder {
	return &MockPathBuilder{}
}

func (b *MockPathBuilder) Build(searchDir, outputDir, interfacePath, interfaceName string) string {
	relativePath, _ := filepath.Rel(searchDir, filepath.Dir(interfacePath))

	mockSubPath := ""
	for _, segment := range strings.Split(relativePath, string(filepath.Separator)) {
		mockSubPath = filepath.Join(mockSubPath, segment+"_mocks")
	}

	fileName := strings.ReplaceAll(utils.ToSnakeCase(interfaceName), "_interface", "")

	return filepath.Join(outputDir, mockSubPath, fmt.Sprintf("%s_mock.go", fileName))
}
