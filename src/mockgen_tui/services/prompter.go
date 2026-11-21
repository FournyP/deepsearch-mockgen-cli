package services

import (
	"errors"
	"strings"

	tea "charm.land/bubbletea/v2"
	tui_models "github.com/FournyP/deepsearch-mockgen-cli/src/mockgen_tui/models"
)

type Prompter struct{}

func NewPrompter() *Prompter {
	return &Prompter{}
}

func (p *Prompter) Prompt(question, defaultValue string) (string, error) {
	final, err := tea.NewProgram(tui_models.NewTextInputModel(question, defaultValue)).Run()
	if err != nil {
		return "", err
	}

	model, ok := final.(tui_models.TextInputModel)
	if !ok {
		return defaultValue, nil
	}
	if model.Cancelled {
		return "", errors.New("cancelled")
	}

	value := strings.TrimSpace(model.Value())
	if value == "" {
		return defaultValue, nil
	}
	return value, nil
}
