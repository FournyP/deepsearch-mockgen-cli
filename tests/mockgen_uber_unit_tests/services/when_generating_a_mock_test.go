package services_test

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/models"
	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen_uber/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// fakeMockgen records its arguments, one per line, then exits with the given code.
const fakeMockgen = "#!/bin/sh\nprintf '%%s\\n' \"$@\" > '%s'\nexit %s\n"

type WhenGeneratingAMockTestingSuite struct {
	sut *services.MockGenerator

	root     string
	argsFile string
}

func WhenGeneratingAMockBeforeEach(t *testing.T, exitCode string) *WhenGeneratingAMockTestingSuite {
	if runtime.GOOS == "windows" {
		t.Skip("fake mockgen is a shell script")
	}

	root := t.TempDir()
	binDir := filepath.Join(root, "bin")
	argsFile := filepath.Join(root, "args")
	require.NoError(t, os.MkdirAll(binDir, 0o755))
	script := fmt.Sprintf(fakeMockgen, argsFile, exitCode)
	require.NoError(t, os.WriteFile(filepath.Join(binDir, "mockgen"), []byte(script), 0o755))
	t.Setenv("PATH", binDir+string(os.PathListSeparator)+os.Getenv("PATH"))

	return &WhenGeneratingAMockTestingSuite{sut: services.NewMockGenerator(), root: root, argsFile: argsFile}
}

func (s *WhenGeneratingAMockTestingSuite) recordedArgs(t *testing.T) []string {
	content, err := os.ReadFile(s.argsFile)
	require.NoError(t, err)
	return strings.Fields(string(content))
}

func TestWhenGeneratingAMock(t *testing.T) {
	t.Run("Given a mock path inside a missing directory", func(t *testing.T) {
		t.Run("Should create the directory and name the package after it", func(t *testing.T) {
			suite := WhenGeneratingAMockBeforeEach(t, "0")
			mockPath := suite.root + "/out/pkg_mocks/saver_mock.go"

			err := suite.sut.Generate(models.MockTarget{
				InterfaceName: "Saver",
				SourcePath:    "src/pkg/saver.go",
				MockPath:      mockPath,
			})

			require.NoError(t, err)
			assert.DirExists(t, filepath.Dir(mockPath))
			assert.Equal(t, []string{
				"-source=src/pkg/saver.go",
				"-destination=" + mockPath,
				"-package=pkg_mocks",
			}, suite.recordedArgs(t))
		})
	})

	t.Run("Given a failing mockgen run", func(t *testing.T) {
		t.Run("Should wrap the failure", func(t *testing.T) {
			suite := WhenGeneratingAMockBeforeEach(t, "1")

			err := suite.sut.Generate(models.MockTarget{
				InterfaceName: "Saver",
				SourcePath:    "src/pkg/saver.go",
				MockPath:      suite.root + "/out/saver_mock.go",
			})

			assert.ErrorContains(t, err, "mockgen execution failed")
		})
	})
}
