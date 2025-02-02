package models

import (
	"context"
	"mathing/internal/store"

	tea "github.com/charmbracelet/bubbletea"
)

func (i *ItemsList) CurrentItem() store.Item {
	return i.items[i.selected]
}

func (i *ItemsList) Refetch() (tea.Model, tea.Cmd) {
	ctx := context.Background()
	var err error
	_, i.data, err = i.store.GetItemTable(ctx, i.pageOffset)
	if err != nil {
		return i, tea.Println(err)
	}

	i.items, err = i.store.ListItems(ctx, i.pageOffset)
	if err != nil {
		return i, tea.Println(err)
	}

	i.itemCount, err = i.store.CountItems(ctx)
	if err != nil {
		return i, tea.Println(err)
	}

	i.selected = 0
	return i, nil
}
