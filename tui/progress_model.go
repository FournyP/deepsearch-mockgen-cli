package tui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

// ProgressUpdate is sent for each finished task.
type ProgressUpdate struct {
	Name string
	Err  error
}

type progressModel struct {
	bar     progress.Model
	total   int
	done    int
	current string
	lines   []string
	updates <-chan ProgressUpdate
}

type updateMsg ProgressUpdate
type updatesDoneMsg struct{}

func NewProgressModel(total int, updates <-chan ProgressUpdate) progressModel {
	pm := progress.New(progress.WithDefaultGradient())
	return progressModel{bar: pm, total: total, updates: updates}
}

func (m progressModel) Init() tea.Cmd {
	return func() tea.Msg {
		// listen for updates and forward as tea messages
		for u := range m.updates {
			return updateMsg(u)
		}
		return updatesDoneMsg{}
	}
}

func (m progressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch fm := msg.(type) {
	case progress.FrameMsg:
		nb, cmd := m.bar.Update(fm)
		if bm, ok := nb.(progress.Model); ok {
			m.bar = bm
		}
		if m.done >= m.total && !m.bar.IsAnimating() {
			return m, tea.Quit
		}
		return m, cmd
	}

	switch um := msg.(type) {
	case updateMsg:
		u := ProgressUpdate(um)
		m.done++
		m.current = u.Name
		if u.Err != nil {
			m.lines = append(m.lines, fmt.Sprintf("%s: failed (%v)", u.Name, u.Err))
		}

		pct := float64(m.done) / float64(m.total)
		cmdAnim := (&m.bar).SetPercent(pct)
		// continue listening for more updates
		listenCmd := func() tea.Msg {
			for u := range m.updates {
				return updateMsg(u)
			}
			return updatesDoneMsg{}
		}
		return m, tea.Batch(cmdAnim, listenCmd)
	case updatesDoneMsg:
		// channel closed. If animation still running, wait for it to finish via FrameMsg.
		if m.bar.IsAnimating() {
			return m, nil
		}
		return m, tea.Quit
	}
	return m, nil
}

func (m progressModel) View() string {
	s := "Generating mocks...\n\n"
	s += m.bar.View() + "\n\n"
	if m.current != "" {
		s += fmt.Sprintf("Current: %s\n\n", m.current)
	}
	for _, l := range m.lines {
		s += l + "\n"
	}
	return s
}

// RunProgress runs the TUI progress bar until the updates channel is closed.
func RunProgress(total int, updates <-chan ProgressUpdate) error {
	m := NewProgressModel(total, updates)
	p := tea.NewProgram(m)
	_, err := p.Run()
	return err
}
