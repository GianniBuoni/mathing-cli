package models

import (
	"context"
	"mathing/internal/lib"

	tea "github.com/charmbracelet/bubbletea"
)

func (i *ItemsList) Delete() (tea.Model, tea.Cmd) {
	if lib.Confirm("Delete item", "Cancel") {
		if err := i.store.DeleteItem(
			context.Background(),
			i.CurrentItem().ID,
		); err != nil {
			return i, tea.Println(err)
		}
	}
	return i.Refetch()
}

func (i *ItemsList) Create() (tea.Model, tea.Cmd) {
	if err := lib.NewItemLoop(i.store); err != nil {
		return i, tea.Println(err)
	}
	return i.Refetch()
}
