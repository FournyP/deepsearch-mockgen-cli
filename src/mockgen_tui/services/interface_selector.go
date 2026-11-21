package services

import (
	"errors"
	"sort"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	tui_models "github.com/FournyP/deepsearch-mockgen-cli/src/mockgen_tui/models"
)

type InterfaceSelector struct{}

func NewInterfaceSelector() *InterfaceSelector {
	return &InterfaceSelector{}
}

func (s *InterfaceSelector) Select(interfaces map[string]string) ([]string, error) {
	names := make([]string, 0, len(interfaces))
	for name := range interfaces {
		names = append(names, name)
	}
	sort.Strings(names)

	items := make([]list.Item, 0, len(names))
	for _, name := range names {
		items = append(items, tui_models.InterfaceItem{Name: name, Path: interfaces[name]})
	}

	final, err := tea.NewProgram(tui_models.NewInterfaceSelectorModel(items)).Run()
	if err != nil {
		return nil, err
	}

	model, ok := final.(tui_models.InterfaceSelectorModel)
	if !ok {
		return nil, nil
	}
	if model.Cancelled {
		return nil, errors.New("cancelled")
	}
	return model.SelectedNames(), nil
}
