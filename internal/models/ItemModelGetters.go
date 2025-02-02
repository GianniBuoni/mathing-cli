package models

import (
	"context"
	"mathing/internal/store"

	tea "github.com/charmbracelet/bubbletea"
)

func (i *ItemModel) Refetch() (tea.Model, tea.Cmd) {
	ctx := context.Background()
	var err error
	_, i.table.data, err = i.store.GetItemTable(ctx, i.table.pageOffset)
	if err != nil {
		return i, tea.Println(err)
	}

	i.items, err = i.store.ListItems(ctx, i.table.pageOffset)
	if err != nil {
		return i, tea.Println(err)
	}

	i.table.itemCount, err = i.store.CountItems(ctx)
	if err != nil {
		return i, tea.Println(err)
	}

	if i.table.selected > len(i.table.data)-1 {
		i.table.selected = len(i.table.data) - 1
	}
	return i, nil
}

func (i *ItemModel) CurrentItem() store.Item {
	return i.items[i.table.selected]
}
