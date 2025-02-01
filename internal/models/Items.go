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
	state      interfaces.State
}

func NewItemsList(s interfaces.State) (*ItemsList, error) {
	ctx := context.Background()
	headers, data, err := s.GetItemTable(ctx, 0)
	if err != nil {
		return nil, err
	}

	count, err := s.GetItemCount(ctx)
	if err != nil {
		return nil, err
	}

	return &ItemsList{
		selected:   0,
		itemCount:  count,
		pageOffset: 0,
		headers:    headers,
		data:       data,
		state:      s,
	}, nil
}

func (i *ItemsList) Init() tea.Cmd {
	return nil
}
