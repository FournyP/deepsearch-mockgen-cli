package generator

import (
	"path/filepath"
	"testing"
)

func TestToSnakeCase(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"SaveLLM", "save_llm"},
		{"HTTPClient", "http_client"},
		{"URLParser", "url_parser"},
		{"Save", "save"},
		{"LLM", "llm"},
		{"UserID", "user_id"},
		{"ID", "id"},
		{"", ""},
		{"SaveLLMInterface", "save_llm_interface"},
	}
	for _, c := range cases {
		if got := toSnakeCase(c.in); got != c.want {
			t.Errorf("toSnakeCase(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestComputeMockPath_AcronymInterface(t *testing.T) {
	searchDir := filepath.Join("/tmp", "project")
	outputDir := filepath.Join("/tmp", "mocks")
	ifacePath := filepath.Join(searchDir, "pkg", "llm.go")

	got := ComputeMockPath(searchDir, outputDir, ifacePath, "SaveLLM")
	want := filepath.Join(outputDir, "pkg_mocks", "save_llm_mock.go")
	if got != want {
		t.Errorf("ComputeMockPath = %q, want %q", got, want)
	}
}

func TestComputeMockPath_StripsInterfaceSuffix(t *testing.T) {
	searchDir := filepath.Join("/tmp", "project")
	outputDir := filepath.Join("/tmp", "mocks")
	ifacePath := filepath.Join(searchDir, "pkg", "llm.go")

	got := ComputeMockPath(searchDir, outputDir, ifacePath, "SaveLLMInterface")
	want := filepath.Join(outputDir, "pkg_mocks", "save_llm_mock.go")
	if got != want {
		t.Errorf("ComputeMockPath = %q, want %q", got, want)
	}
}
