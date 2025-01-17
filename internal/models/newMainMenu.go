package models

import tea "github.com/charmbracelet/bubbletea"

func newMainMenu(c *config) tea.Model {
	return menu{
		choices:  []string{"Update receipt", "New receipt", "Update item", "New item"},
		selected: 0,
		config:   c,
		name:     mainMenu,
	}
}
