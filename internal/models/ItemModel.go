package models

import (
	"context"
	"mathing/internal/store"

	tea "github.com/charmbracelet/bubbletea"
)

func NewItemModel(s *store.ItemStore) (*ListModel[store.Item], error) {
	lm := &ListModel[store.Item]{
		store:       s,
		actionFuncs: map[ListAction]func() tea.Cmd{},
	}
	lm.table = NewTableData()
	lm.RegisterAction(remove, lm.Delete)
	lm.RegisterAction(create, lm.Post)
	lm.RegisterAction(edit, lm.Post)

	ctx := context.Background()
	var err error
	lm.table.headers, lm.table.data, err = s.GetTable(ctx, 0)
	if err != nil {
		return nil, err
	}
	lm.table.itemCount, err = s.CountRows(ctx)
	if err != nil {
		return nil, err
	}
	lm.rows, err = s.GetRows(ctx, 0)
	if err != nil {
		return nil, err
	}

	return lm, nil
}
