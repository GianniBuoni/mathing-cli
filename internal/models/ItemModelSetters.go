package models

import (
	"context"
	"mathing/internal/lib"

	tea "github.com/charmbracelet/bubbletea"
)

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

func (i *ItemModel) Edit() tea.Cmd {
	res, err := lib.NewItemParser(i.form)
	if err != nil {
		return tea.Println(err)
	}

	res.ID = i.CurrentItem().ID

	err = i.store.CreateItem(context.Background(), res)
	if err != nil {
		return tea.Println(err)
	}
	if err := i.Refetch(); err != nil {
		return tea.Println(err)
	}
	return nil
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
