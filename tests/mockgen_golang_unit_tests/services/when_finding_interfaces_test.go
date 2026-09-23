package services_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen_golang/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type WhenFindingInterfacesTestingSuite struct {
	sut *services.InterfaceFinder

	root string
}

func WhenFindingInterfacesBeforeEach(t *testing.T) *WhenFindingInterfacesTestingSuite {
	return &WhenFindingInterfacesTestingSuite{sut: services.NewInterfaceFinder(), root: t.TempDir()}
}

func (s *WhenFindingInterfacesTestingSuite) writeFile(t *testing.T, relativePath, content string) string {
	path := filepath.Join(s.root, relativePath)
	require.NoError(t, os.MkdirAll(filepath.Dir(path), 0o755))
	require.NoError(t, os.WriteFile(path, []byte(content), 0o644))
	return path
}

func TestWhenFindingInterfaces(t *testing.T) {
	t.Parallel()

	t.Run("Given a tree of Go files", func(t *testing.T) {
		t.Parallel()

		t.Run("Should collect every interface with its source file", func(t *testing.T) {
			t.Parallel()

			suite := WhenFindingInterfacesBeforeEach(t)
			first := suite.writeFile(t, "pkg/llm.go", "package pkg\ntype Saver interface{ Save() }\ntype Model struct{}\n")
			second := suite.writeFile(t, "pkg/sub/http.go", "package sub\ntype (\n\tHTTPClient interface{ Do() }\n\tReader interface{ Read() }\n)\n")

			found, err := suite.sut.Find(suite.root)

			require.NoError(t, err)
			assert.Equal(t, map[string]string{
				"Saver":      first,
				"HTTPClient": second,
				"Reader":     second,
			}, found)
		})

		t.Run("Should skip test, non-Go and unparsable files", func(t *testing.T) {
			t.Parallel()

			suite := WhenFindingInterfacesBeforeEach(t)
			suite.writeFile(t, "pkg/llm_test.go", "package pkg\ntype Tested interface{ T() }\n")
			suite.writeFile(t, "pkg/notes.txt", "type Text interface{}\n")
			suite.writeFile(t, "pkg/broken.go", "package pkg\ntype Broken interface{\n")

			found, err := suite.sut.Find(suite.root)

			require.NoError(t, err)
			assert.Empty(t, found)
		})
	})

	t.Run("Given a missing search directory", func(t *testing.T) {
		t.Parallel()

		t.Run("Should find nothing without failing", func(t *testing.T) {
			t.Parallel()

			suite := WhenFindingInterfacesBeforeEach(t)

			found, err := suite.sut.Find(filepath.Join(suite.root, "missing"))

			require.NoError(t, err)
			assert.Empty(t, found)
		})
	})
}
