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

func UpdateGenerator() func(tea.Msg, Table) (tea.Model, tea.Cmd) {
	return func(msg tea.Msg, table Table) (tea.Model, tea.Cmd) {
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
				return table.Refetch()
			case "h", "left":
				table.PagePrev()
				return table.Refetch()
			case "a":
				return table.Create()
			case "d":
				return table.Delete()
			case "ctrl+c":
				return table, tea.Quit
			}
		}
		return table, nil
	}
}
