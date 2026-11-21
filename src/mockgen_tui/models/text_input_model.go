package models

import (
	"fmt"

	textinput "charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

type TextInputModel struct {
	input     textinput.Model
	question  string
	Cancelled bool
}

func NewTextInputModel(question, defaultValue string) TextInputModel {
	input := textinput.New()
	input.Placeholder = defaultValue
	input.SetValue(defaultValue)
	input.Focus()

	return TextInputModel{input: input, question: question}
}

func (m TextInputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m TextInputModel) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	var command tea.Cmd
	m.input, command = m.input.Update(message)

	if typed, ok := message.(tea.KeyPressMsg); ok {
		if typed.String() == "enter" {
			return m, tea.Quit
		}
		if typed.String() == "ctrl+c" {
			m.Cancelled = true
			return m, tea.Quit
		}
	}

	return m, command
}

func (m TextInputModel) View() tea.View {
	return tea.NewView(fmt.Sprintf("%s\n\n%s\n\n(Enter to confirm)", m.question, m.input.View()))
}

func (m TextInputModel) Value() string {
	return m.input.Value()
}
