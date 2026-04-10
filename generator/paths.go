package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

// ComputeMockPath generates a default output path for a mock file.
func ComputeMockPath(searchDir, outputDir, ifacePath, ifaceName string) string {
	// Get relative path from searchDir
	relPath, _ := filepath.Rel(searchDir, filepath.Dir(ifacePath))

	mockSubPath := ""
	for _, s := range strings.Split(relPath, "/") {
		// Add _mocks to the path
		mockSubPath = filepath.Join(mockSubPath, s+"_mocks")
	}

	mockDir := filepath.Join(outputDir, mockSubPath)

	// Convert ifaceName from PascalCase to snake_case
	snakeCaseIfaceName := toSnakeCase(ifaceName)
	snakeCaseIfaceName = strings.ReplaceAll(snakeCaseIfaceName, "_interface", "")

	// Define mock filename
	mockFile := fmt.Sprintf("%s_mock.go", snakeCaseIfaceName)

	return filepath.Join(mockDir, mockFile)
}

// toSnakeCase converts a PascalCase/CamelCase string to snake_case,
// treating runs of uppercase letters as acronyms (e.g. "SaveLLM" -> "save_llm",
// "HTTPClient" -> "http_client", "URLParser" -> "url_parser").
func toSnakeCase(str string) string {
	runes := []rune(str)
	var result []rune
	for i, r := range runes {
		if unicode.IsUpper(r) {
			if i > 0 {
				prev := runes[i-1]
				prevIsLowerOrDigit := unicode.IsLower(prev) || unicode.IsDigit(prev)
				nextIsLower := i+1 < len(runes) && unicode.IsLower(runes[i+1])
				if prevIsLowerOrDigit || (unicode.IsUpper(prev) && nextIsLower) {
					result = append(result, '_')
				}
			}
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

func CreateDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}
	return nil
}
