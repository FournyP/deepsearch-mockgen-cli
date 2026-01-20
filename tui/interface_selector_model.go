package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type interfaceSelectorModel struct {
	lst       list.Model
	selected  map[string]bool
	done      bool
	Cancelled bool
}

func NewInterfaceSelector(items []list.Item) interfaceSelectorModel {
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Interfaces"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	return interfaceSelectorModel{lst: l, selected: make(map[string]bool), done: false}
}

func (m interfaceSelectorModel) Init() tea.Cmd { return nil }

func (m interfaceSelectorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	newLst, c := m.lst.Update(msg)
	m.lst = newLst
	cmd = c
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		reserved := 6
		w := msg.Width
		h := msg.Height - reserved
		if h < 3 {
			h = 3
		}
		m.lst.SetSize(w, h)
		return m, cmd

	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			m.done = true
			return m, tea.Quit
		case "ctrl+c":
			m.done = true
			m.Cancelled = true
			return m, tea.Quit
		default:
			// also respond to explicit Ctrl+C key type
			if msg.Type == tea.KeyCtrlC {
				m.done = true
				m.Cancelled = true
				return m, tea.Quit
			}
		case " ":
			idx := m.lst.Index()
			if idx >= 0 && idx < len(m.lst.Items()) {
				if it, ok := m.lst.Items()[idx].(ifaceItem); ok {
					it.Selected = !it.Selected
					m.lst.SetItem(idx, it)
					m.selected[it.Name] = it.Selected
				}
			}
		case "enter":
			m.done = true
			return m, tea.Quit
		}
	}

	return m, cmd
}

func (m interfaceSelectorModel) View() string {
	if m.done {
		return "Generating mocks...\n"
	}
	s := "Select interfaces to generate mocks for:\n\n"
	s += m.lst.View()
	s += "\n\n[Space] toggle | [Enter] confirm selection | [Q] quit\n"
	return s
}

func (m interfaceSelectorModel) SelectedMap() map[string]bool { return m.selected }

type ifaceItem struct {
	Name     string
	Path     string
	Selected bool
}

func (i ifaceItem) Title() string {
	if i.Selected {
		return fmt.Sprintf("[x] %s", i.Name)
	}
	return fmt.Sprintf("[ ] %s", i.Name)
}

func (i ifaceItem) Description() string { return i.Path }

func (i ifaceItem) FilterValue() string { return i.Name }
