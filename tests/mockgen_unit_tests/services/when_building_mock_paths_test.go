package services_test

import (
	"path/filepath"
	"testing"

	"github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/services"
	"github.com/stretchr/testify/assert"
)

type WhenBuildingMockPathsTestingSuite struct {
	sut *services.MockPathBuilder

	searchDir string
	outputDir string
}

func WhenBuildingMockPathsBeforeEach(_ *testing.T) *WhenBuildingMockPathsTestingSuite {
	return &WhenBuildingMockPathsTestingSuite{
		sut:       services.NewMockPathBuilder(),
		searchDir: filepath.Join("/tmp", "project"),
		outputDir: filepath.Join("/tmp", "mocks"),
	}
}

func TestWhenBuildingMockPaths(t *testing.T) {
	t.Parallel()

	t.Run("Given an interface named with an acronym", func(t *testing.T) {
		t.Parallel()

		t.Run("Should snake case the mock file name", func(t *testing.T) {
			t.Parallel()

			suite := WhenBuildingMockPathsBeforeEach(t)
			interfacePath := filepath.Join(suite.searchDir, "pkg", "llm.go")

			assert.Equal(
				t,
				filepath.Join(suite.outputDir, "pkg_mocks", "save_llm_mock.go"),
				suite.sut.Build(suite.searchDir, suite.outputDir, interfacePath, "SaveLLM"),
			)
		})
	})

	t.Run("Given an interface suffixed with Interface", func(t *testing.T) {
		t.Parallel()

		t.Run("Should strip the suffix from the mock file name", func(t *testing.T) {
			t.Parallel()

			suite := WhenBuildingMockPathsBeforeEach(t)
			interfacePath := filepath.Join(suite.searchDir, "pkg", "llm.go")

			assert.Equal(
				t,
				filepath.Join(suite.outputDir, "pkg_mocks", "save_llm_mock.go"),
				suite.sut.Build(suite.searchDir, suite.outputDir, interfacePath, "SaveLLMInterface"),
			)
		})
	})

	t.Run("Given an interface nested several directories deep", func(t *testing.T) {
		t.Parallel()

		t.Run("Should suffix every mirrored directory with _mocks", func(t *testing.T) {
			t.Parallel()

			suite := WhenBuildingMockPathsBeforeEach(t)
			interfacePath := filepath.Join(suite.searchDir, "services", "interfaces", "prompter_interface.go")

			assert.Equal(
				t,
				filepath.Join(suite.outputDir, "services_mocks", "interfaces_mocks", "prompter_mock.go"),
				suite.sut.Build(suite.searchDir, suite.outputDir, interfacePath, "PrompterInterface"),
			)
		})
	})
}
