package models

import "fmt"

type InterfaceItem struct {
	Name     string
	Path     string
	Selected bool
}

func (i InterfaceItem) Title() string {
	if i.Selected {
		return fmt.Sprintf("[x] %s", i.Name)
	}
	return fmt.Sprintf("[ ] %s", i.Name)
}

func (i InterfaceItem) Description() string {
	return i.Path
}

func (i InterfaceItem) FilterValue() string {
	return i.Name
}
