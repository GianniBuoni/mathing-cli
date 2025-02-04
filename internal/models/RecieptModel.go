package models

import (
	"context"
	"mathing/internal/store"

	tea "github.com/charmbracelet/bubbletea"
)

func NewRecieptModel(s *store.RecieptStore) (
	*ListModel[store.ListReceiptRow], error,
) {
	lm := &ListModel[store.ListReceiptRow]{
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
