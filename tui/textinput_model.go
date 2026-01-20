package tui

import (
	"fmt"

	textinput "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type textInputModel struct {
	ti       textinput.Model
	question string
}

func (m textInputModel) Init() tea.Cmd { return textinput.Blink }

func (m textInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.ti, cmd = m.ti.Update(msg)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "enter" {
			return m, tea.Quit
		}
	}
	return m, cmd
}

func (m textInputModel) View() string {
	return fmt.Sprintf("%s\n\n%s\n\n(Enter to confirm)", m.question, m.ti.View())
}
