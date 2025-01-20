package models

import (
	"context"
	"fmt"
	"mathing/internal/store"
)

type ItemList struct {
	tableMenu
}

func NewItemList(s *store.Queries) (subModel, error) {
	list := &ItemList{
		tableMenu{
			state:       listItems,
			store:       s,
			selected:    0,
			offset:      0,
			offsetSteps: 20,
			headers:     []string{"ITEM NAME", "PRICE"},
		},
	}

	if err := list.Get(); err != nil {
		return nil, err
	}
	return list, nil
}

func (i *ItemList) Get() error {
	ctx := context.Background()
	res, err := i.store.ListItems(ctx, int64(i.offset))
	if err != nil {
		return fmt.Errorf("issue with getting item list: %w", err)
	}

	count, err := i.store.CountItems(ctx)
	if err != nil {
		return fmt.Errorf("issue getting item count: %w", err)
	}

	itemIds := []int{}
	content := [][]string{}

	for _, v := range res {
		row := []string{v.Item, fmt.Sprintf("%05.2f", v.Price)}
		content = append(content, row)

		itemIds = append(itemIds, int(v.ID))
	}

	// update values
	i.itemCount = int(count)
	i.itemIDs = itemIds
	i.content = content

	return nil
}

func (i *ItemList) Upsert() error {
	return nil
}

func (i *ItemList) Delete() error {
	return nil
}

func (i *ItemList) NextState() state {
	return mainMenu
}
