package commands

import (
	"context"
	"fmt"
)

func (s *State) GetItemTable(ctx context.Context) (
	headers []string, data [][]string, err error,
) {
	items, err := s.Store.ListItems(ctx, 0)
	if err != nil {
		return nil, nil, fmt.Errorf("issue getting item data: %w", err)
	}

	for _, v := range items {
		price := fmt.Sprintf("%05.2f", v.Price)
		row := []string{v.Item, price}
		data = append(data, row)
	}

	headers = []string{"NAME", "PRICE"}

	return headers, data, nil
}

func (s *State) GetUserTable(ctx context.Context) (
	headers []string, data [][]string, err error,
) {
	users, err := s.Store.ListUsers(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("issue getting user data: %w", err)
	}

	for _, v := range users {
		row := []string{v.Name}
		data = append(data, row)
	}

	headers = []string{"NAME"}

	return headers, data, nil
}
