package models

import (
	"context"
	"fmt"
	"mathing/internal/store"
)

type ListUsers struct {
	tableMenu
}

// display as list instead of a table?
func NewListUsers(s *store.Queries) (subModel, error) {
	// hard coded data members
	menu := &ListUsers{
		tableMenu{
			state:       listUsers,
			store:       s,
			selected:    0,
			offset:      0,
			offsetSteps: 10,
			headers:     []string{"NAME"},
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

	content := [][]string{}
	itemIDs := []int{}

	for _, v := range res {
		row := []string{v.Name}
		content = append(content, row)
		itemIDs = append(itemIDs, int(v.ID))
	}

	count := len(content)

	l.content = content
	l.itemIDs = itemIDs
	l.itemCount = count

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
