package models

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

const reservedHeight = 6

type InterfaceSelectorModel struct {
	list      list.Model
	done      bool
	Cancelled bool
}

func NewInterfaceSelectorModel(items []list.Item) InterfaceSelectorModel {
	interfaceList := list.New(items, list.NewDefaultDelegate(), 0, 0)
	interfaceList.Title = "Interfaces"
	interfaceList.SetShowStatusBar(false)
	interfaceList.SetFilteringEnabled(false)

	return InterfaceSelectorModel{list: interfaceList}
}

func (m InterfaceSelectorModel) Init() tea.Cmd {
	return nil
}

func (m InterfaceSelectorModel) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	var command tea.Cmd
	m.list, command = m.list.Update(message)

	switch typed := message.(type) {
	case tea.WindowSizeMsg:
		m.list.SetSize(typed.Width, max(typed.Height-reservedHeight, 3))
	case tea.KeyMsg:
		switch {
		case typed.Type == tea.KeyCtrlC || typed.String() == "ctrl+c":
			m.done = true
			m.Cancelled = true
			return m, tea.Quit
		case typed.String() == "q", typed.String() == "enter":
			m.done = true
			return m, tea.Quit
		case typed.String() == " ":
			m.toggleCurrent()
		}
	}

	return m, command
}

func (m InterfaceSelectorModel) View() string {
	if m.done {
		return "Generating mocks...\n"
	}
	return "Select interfaces to generate mocks for:\n\n" +
		m.list.View() +
		"\n\n[Space] toggle | [Enter] confirm selection | [Q] quit\n"
}

// SelectedNames returns the names of the toggled interfaces.
func (m InterfaceSelectorModel) SelectedNames() []string {
	names := make([]string, 0)
	for _, item := range m.list.Items() {
		if typed, ok := item.(InterfaceItem); ok && typed.Selected {
			names = append(names, typed.Name)
		}
	}
	return names
}

func (m *InterfaceSelectorModel) toggleCurrent() {
	index := m.list.Index()
	if index < 0 || index >= len(m.list.Items()) {
		return
	}
	if item, ok := m.list.Items()[index].(InterfaceItem); ok {
		item.Selected = !item.Selected
		m.list.SetItem(index, item)
	}
}
