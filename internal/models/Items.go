package models

import (
	"context"
	"mathing/internal/interfaces"
	"mathing/internal/store"

	tea "github.com/charmbracelet/bubbletea"
)

type ItemsList struct {
	TableData
	items  []store.Item
	update func(tea.Msg, Table) (tea.Model, tea.Cmd)
}

func NewItemsList(s interfaces.Store) (*ItemsList, error) {
	ctx := context.Background()
	headers, data, err := s.GetItemTable(ctx, 0)
	if err != nil {
		return nil, err
	}

	items, err := s.ListItems(ctx, 0)
	if err != nil {
		return nil, err
	}

	count, err := s.CountItems(ctx)
	if err != nil {
		return nil, err
	}

	return &ItemsList{
		TableData: TableData{
			selected:   0,
			itemCount:  count,
			pageOffset: 0,
			store:      s,
			headers:    headers,
			data:       data,
		},
		items:  items,
		update: HadnleUpdate,
	}, nil
}
