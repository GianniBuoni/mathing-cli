package models

import (
	"context"
	"fmt"
	"mathing/internal/store"

	tea "github.com/charmbracelet/bubbletea"
)

func (i *ItemsList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if i.update == nil {
		fmt.Println("❗Update method not implemented: quitting program")
		return i, tea.Quit
	}
	return i.update(msg, i)
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

	if i.selected > len(i.data)-1 {
		i.selected = len(i.data) - 1
		return i, tea.Println(i.selected)
	}
	return i, nil
}

func (i *ItemsList) CurrentItem() store.Item {
	return i.items[i.selected]
}
