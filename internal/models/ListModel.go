package models

import (
	"context"
	"mathing/internal/interfaces"
	"mathing/internal/store"

	"github.com/charmbracelet/huh"
)

type ListState uint

const (
	table ListState = iota
	form
)

type ListModel struct {
	state ListState
	table *TableData
	form  *huh.Form
	store interfaces.Store
}

type ItemModel struct {
	items []store.Item
  ListModel
}

func NewItemsList(s interfaces.Store) (*ItemModel, error) {
	lm := &ItemModel{
    ListModel: ListModel{
      state: table,
      store: s,
    },
	}
	lm.table = NewTableData()

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
