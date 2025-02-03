package models

import (
	"context"
	"fmt"
	"mathing/internal/lib"

	tea "github.com/charmbracelet/bubbletea"
)

// CREATE
func (i *ItemModel) CreateInit() tea.Cmd {
	i.form = lib.NewItemForm()
	i.state = form
	i.action = create
	return i.form.Init()
}

func (i *ItemModel) Create() tea.Cmd {
	res, err := lib.NewItemParser(i.form)
	if err != nil {
		return tea.Println(err)
	}
	err = i.store.CreateItem(context.Background(), res)
	if err != nil {
		return tea.Println(err)
	}
	if err := i.Refetch(); err != nil {
		return tea.Println(err)
	}
	return nil
}

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
