package models

import tea "github.com/charmbracelet/bubbletea"

func newReceipt(c *config) tea.Model {
	return menu{
		choices:  []string{"taquitos", "samosas", "popcorn", "pretzels"},
		selected: 0,
		config:   c,
		name:     receipt,
	}
}
