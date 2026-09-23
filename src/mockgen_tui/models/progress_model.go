package models

import (
	"fmt"

	mockgen_models "github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/models"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type progressUpdateMsg mockgen_models.ProgressUpdate

type progressDoneMsg struct{}

type ProgressModel struct {
	bar     progress.Model
	total   int
	done    int
	current string
	lines   []string
	updates <-chan mockgen_models.ProgressUpdate
}

func NewProgressModel(total int, updates <-chan mockgen_models.ProgressUpdate) ProgressModel {
	return ProgressModel{
		bar:     progress.New(progress.WithDefaultGradient()),
		total:   total,
		updates: updates,
	}
}

func (m ProgressModel) Init() tea.Cmd {
	return m.listen()
}

func (m ProgressModel) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch typed := message.(type) {
	case progress.FrameMsg:
		nextBar, command := m.bar.Update(typed)
		if bar, ok := nextBar.(progress.Model); ok {
			m.bar = bar
		}
		if m.done >= m.total && !m.bar.IsAnimating() {
			return m, tea.Quit
		}
		return m, command
	case progressUpdateMsg:
		update := mockgen_models.ProgressUpdate(typed)
		m.done++
		m.current = update.Name
		if update.Err != nil {
			m.lines = append(m.lines, fmt.Sprintf("%s: failed (%v)", update.Name, update.Err))
		}
		animation := (&m.bar).SetPercent(float64(m.done) / float64(m.total))
		return m, tea.Batch(animation, m.listen())
	case progressDoneMsg:
		if m.bar.IsAnimating() {
			return m, nil
		}
		return m, tea.Quit
	}

	return m, nil
}

func (m ProgressModel) View() string {
	view := "Generating mocks...\n\n" + m.bar.View() + "\n\n"
	if m.current != "" {
		view += fmt.Sprintf("Current: %s\n\n", m.current)
	}
	for _, line := range m.lines {
		view += line + "\n"
	}
	return view
}

func (m ProgressModel) listen() tea.Cmd {
	return func() tea.Msg {
		for update := range m.updates {
			return progressUpdateMsg(update)
		}
		return progressDoneMsg{}
	}
}
