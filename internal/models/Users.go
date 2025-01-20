package models

import (
	"context"
	"fmt"
	"mathing/internal/store"
)

type ListUsers struct {
	listMenu
}

// display as list instead of a table?
func NewListUsers(s *store.Queries) (subModel, error) {
	// hard coded data members
	menu := &ListUsers{
		listMenu{
			state:    listUsers,
			store:    s,
			selected: 0,
		},
	}

	// fetched data
	if err := menu.Get(); err != nil {
		return nil, err
	}
	return menu, nil
}

func (l *ListUsers) Get() error {
	cxt := context.Background()
	res, err := l.store.ListUsers(cxt)
	if err != nil {
		return fmt.Errorf("issue fetching user data: %w", err)
	}

	content := []string{}
	itemIDs := []int{}

	for _, v := range res {
		content = append(content, v.Name)
		itemIDs = append(itemIDs, int(v.ID))
	}

	l.content = content
	l.itemIDs = itemIDs

	return nil
}

func (l *ListUsers) Upsert() error {
	return nil
}

func (l *ListUsers) Delete() error {
	return nil
}

func (l *ListUsers) NextState() state {
	return mainMenu
}
