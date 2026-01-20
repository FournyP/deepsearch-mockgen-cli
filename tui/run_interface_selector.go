package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func RunInterfaceSelector(interfaces map[string]string, outputDir string) (map[string]string, error) {
	items := make([]list.Item, 0, len(interfaces))
	for k, v := range interfaces {
		items = append(items, ifaceItem{Name: k, Path: v, Selected: false})
	}

	ism := NewInterfaceSelector(items)
	p := tea.NewProgram(ism)
	final, err := p.Run()
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	if fm, ok := final.(interfaceSelectorModel); ok {
		for name, sel := range fm.SelectedMap() {
			if sel {
				result[name] = outputDir
			}
		}
		if len(result) == 0 {
			for _, it := range fm.lst.Items() {
				if ii, ok := it.(ifaceItem); ok {
					if ii.Selected {
						result[ii.Name] = outputDir
					}
				}
			}
		}
	}

	return result, nil
}
