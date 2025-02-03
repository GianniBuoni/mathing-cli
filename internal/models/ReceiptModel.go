package models

import (
	"context"
	"mathing/internal/interfaces"
	"mathing/internal/store"

	tea "github.com/charmbracelet/bubbletea"
)

type ReceiptModel struct {
	receiptRows []store.ListReceiptRow
	ListModel
}

func NewReceiptList(s interfaces.Store) (*ReceiptModel, error) {
	rm := &ReceiptModel{
		ListModel: ListModel{
			store:       s,
			actionFuncs: map[ListAction]func() tea.Cmd{},
		},
	}

	rm.table = NewTableData()
	// TODO register actions

	ctx := context.Background()
	var err error

	rm.table.headers, rm.table.data, err = s.GetReceiptTable(ctx)
	if err != nil {
		return nil, err
	}

	rm.table.itemCount, err = s.CountReceipt(ctx)
	if err != nil {
		return nil, err
	}

	rm.receiptRows, err = s.ListReceipt(ctx)
	if err != nil {
		return nil, err
	}

	return rm, nil
}
