package models

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (table *TableData) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// navigation
		case "j", "down":
			table.SelectNext()
		case "k", "up":
			table.SelectPrev()
		case "l", "right":
			table.PageNext()
		case "h", "left":
			table.PagePrev()
		}
	}
	return table, nil
}
