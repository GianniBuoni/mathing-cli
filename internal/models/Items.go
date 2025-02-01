package models

import (
	"context"
	"mathing/internal/interfaces"

	tea "github.com/charmbracelet/bubbletea"
)

type ItemsList struct {
	selected   int
	itemCount  int64
	pageOffset int64
	headers    []string
	data       [][]string
	store      interfaces.Store
}

func NewItemsList(s interfaces.Store) (*ItemsList, error) {
	ctx := context.Background()
	headers, data, err := s.GetItemTable(ctx, 0)
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
		headers:    headers,
		data:       data,
		store:      s,
	}, nil
}

func (i *ItemsList) Init() tea.Cmd {
	return nil
}
