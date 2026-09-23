package utils_test

import (
	"testing"

	"github.com/FournyP/deepsearch-mockgen-cli/src/common/utils"
	"github.com/stretchr/testify/assert"
)

func TestWhenConvertingToSnakeCase(t *testing.T) {
	t.Parallel()

	t.Run("Given a PascalCase name", func(t *testing.T) {
		t.Parallel()

		t.Run("Should split words with underscores", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, "save", utils.ToSnakeCase("Save"))
			assert.Equal(t, "save_llm_interface", utils.ToSnakeCase("SaveLLMInterface"))
			assert.Equal(t, "", utils.ToSnakeCase(""))
		})
	})

	t.Run("Given a name containing acronyms", func(t *testing.T) {
		t.Parallel()

		t.Run("Should keep each acronym as a single word", func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, "save_llm", utils.ToSnakeCase("SaveLLM"))
			assert.Equal(t, "http_client", utils.ToSnakeCase("HTTPClient"))
			assert.Equal(t, "url_parser", utils.ToSnakeCase("URLParser"))
			assert.Equal(t, "llm", utils.ToSnakeCase("LLM"))
			assert.Equal(t, "user_id", utils.ToSnakeCase("UserID"))
			assert.Equal(t, "id", utils.ToSnakeCase("ID"))
		})
	})
}
