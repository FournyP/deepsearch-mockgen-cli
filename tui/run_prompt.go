package tui

import (
	"strings"

	textinput "github.com/charmbracelet/bubbles/textinput"
)

func RunTextInputPrompt(question, defaultValue string) (string, error) {
	ti := textinput.New()
	ti.Placeholder = defaultValue
	ti.SetValue(defaultValue)
	ti.Focus()

	m := textInputModel{ti: ti, question: question}
	finalM, err := RunTextInputModel(m)
	if err != nil {
		return "", err
	}

	val := strings.TrimSpace(finalM.ti.Value())
	if val == "" {
		return defaultValue, nil
	}
	return val, nil
}
