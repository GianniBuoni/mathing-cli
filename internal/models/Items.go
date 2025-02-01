package models

import (
	"context"
	"mathing/internal/interfaces"
	"mathing/internal/store"

	tea "github.com/charmbracelet/bubbletea"
)

type ItemsList struct {
	selected   int
	itemCount  int64
	pageOffset int64
	store      interfaces.Store
	headers    []string
	data       [][]string
	items      []store.Item
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
		selected:   0,
		itemCount:  count,
		pageOffset: 0,
		store:      s,
		headers:    headers,
		data:       data,
		items:      items,
	}, nil
}

func (i *ItemsList) Init() tea.Cmd {
	return nil
}
