package services

import (
	tea "charm.land/bubbletea/v2"
	mockgen_models "github.com/FournyP/deepsearch-mockgen-cli/src/mockgen/models"
	tui_models "github.com/FournyP/deepsearch-mockgen-cli/src/mockgen_tui/models"
)

type ProgressReporter struct{}

func NewProgressReporter() *ProgressReporter {
	return &ProgressReporter{}
}

func (r *ProgressReporter) Report(total int, updates <-chan mockgen_models.ProgressUpdate) error {
	_, err := tea.NewProgram(tui_models.NewProgressModel(total, updates)).Run()
	return err
}
