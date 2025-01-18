package models

import (
	"context"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func newReceipt(c *config) (tea.Model, error) {
	items, err := c.store.ListItems(context.Background())
	if err != nil {
		return nil, fmt.Errorf("issue retrieving items: %w", err)
	}

	choices := []string{}
	for _, item := range items {
		choices = append(choices, item.Item)
	}

	return menu{
		choices:  choices,
		selected: 0,
		config:   c,
		name:     receipt,
	}, nil
}
