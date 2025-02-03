package models

import (
	"context"
	"mathing/internal/store"

	tea "github.com/charmbracelet/bubbletea"
)

func (i *ItemModel) Refetch() tea.Cmd {
	ctx := context.Background()
	var err error
	// get count first
	i.table.itemCount, err = i.store.CountItems(ctx)
	if err != nil {
		return tea.Println(err)
	}

	// change page offset if needed
	if i.table.CurrentPage() > i.table.PageCount() {
		i.table.pageOffset = (i.table.PageCount() - 1) * 20
	}

	// get new table data
	_, i.table.data, err = i.store.GetItemTable(ctx, i.table.pageOffset)
	if err != nil {
		return tea.Println(err)
	}

	i.items, err = i.store.ListItems(ctx, i.table.pageOffset)
	if err != nil {
		return tea.Println(err)
	}

	if i.table.selected > len(i.table.data)-1 {
		i.table.selected = len(i.table.data) - 1
	}

	return nil
}

func (i *ItemModel) CurrentItem() store.Item {
	return i.items[i.table.selected]
}
