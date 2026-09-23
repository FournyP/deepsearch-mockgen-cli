package utils

import "unicode"

// ToSnakeCase converts a PascalCase/CamelCase string to snake_case,
// treating runs of uppercase letters as acronyms (e.g. "SaveLLM" -> "save_llm",
// "HTTPClient" -> "http_client", "URLParser" -> "url_parser").
func ToSnakeCase(value string) string {
	runes := []rune(value)
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
