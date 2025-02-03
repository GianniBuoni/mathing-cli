package models

import (
	"context"
	"mathing/internal/interfaces"
	"mathing/internal/store"

	tea "github.com/charmbracelet/bubbletea"
)

type ItemModel struct {
	items []store.Item
	ListModel
}

func NewItemsList(s interfaces.Store) (*ItemModel, error) {
	lm := &ItemModel{
		ListModel: ListModel{
			store: s,
      actionFuncs: map[ListAction]func() tea.Cmd{},
		},
	}
	lm.table = NewTableData()
  lm.RegisterAction(remove, lm.Delete)

	ctx := context.Background()
	var err error
	lm.table.headers, lm.table.data, err = s.GetItemTable(ctx, 0)
	if err != nil {
		return nil, err
	}
	lm.table.itemCount, err = s.CountItems(ctx)
	if err != nil {
		return nil, err
	}
	lm.items, err = s.ListItems(ctx, 0)
	if err != nil {
		return nil, err
	}

	return lm, nil
}
