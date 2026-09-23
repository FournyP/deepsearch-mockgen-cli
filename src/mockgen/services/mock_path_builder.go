package services

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/FournyP/deepsearch-mockgen-cli/src/common/utils"
)

const rootMockDir = "mocks"

// MockPathBuilder computes the default output path of a mock: the interface
// source tree relative to searchDir is mirrored under outputDir with every
// directory suffixed by "_mocks". Interfaces declared directly in searchDir
// go to a "mocks" directory, since "._mocks" is not a valid package name.
type MockPathBuilder struct{}

func NewMockPathBuilder() *MockPathBuilder {
	return &MockPathBuilder{}
}

func (b *MockPathBuilder) Build(searchDir, outputDir, interfacePath, interfaceName string) string {
	fileName := strings.ReplaceAll(utils.ToSnakeCase(interfaceName), "_interface", "")

	return filepath.Join(outputDir, mockSubPath(searchDir, interfacePath), fmt.Sprintf("%s_mock.go", fileName))
}

func mockSubPath(searchDir, interfacePath string) string {
	relativePath, _ := filepath.Rel(searchDir, filepath.Dir(interfacePath))
	if relativePath == "." {
		return rootMockDir
	}

	subPath := ""
	for _, segment := range strings.Split(relativePath, string(filepath.Separator)) {
		subPath = filepath.Join(subPath, segment+"_mocks")
	}
	return subPath
}
