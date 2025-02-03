package models

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Table interface {
	tea.Model
	SelectNext()
	SelectPrev()
	PageNext()
	PagePrev()
	Refetch() (tea.Model, tea.Cmd)
	Create() (tea.Model, tea.Cmd)
	Delete() (tea.Model, tea.Cmd)
}

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
