package models

import (
	"context"
	"fmt"
	"mathing/internal/lib"

	tea "github.com/charmbracelet/bubbletea"
)

var confirm string

func (i *ListModel) RegisterAction(la ListAction, f func() tea.Cmd) {
	i.actionFuncs[la] = f
}

func (i *ItemModel) Create() {}

func (i *ItemModel) Edit() {}

// DELETE
func (i *ItemModel) DeleteInit() tea.Cmd {
	title := fmt.Sprintf("Delete %s?", i.CurrentItem().Item)
	i.form = lib.NewDeleteForm(title)
	i.state = form
	i.action = remove
	return i.form.Init()
}

func (i *ItemModel) Delete() tea.Cmd {
	if err := i.store.DeleteItem(
		context.Background(), i.CurrentItem().ID,
	); err != nil {
		return tea.Println(err)
	}
	if err := i.Refetch(); err != nil {
		return tea.Println(err)
	}
	return nil
}
