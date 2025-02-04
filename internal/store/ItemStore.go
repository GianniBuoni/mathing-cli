package store

import (
	"context"
	"fmt"

	"github.com/charmbracelet/huh"
)

type ItemStore struct {
	queries *Queries
}

func NewItemStore(q *Queries) *ItemStore {
	return &ItemStore{
		queries: q,
	}
}

func (i *ItemStore) GetTable(ctx context.Context, pageOffset int64) (
	headers []string, data [][]string, err error,
) {
	items, err := i.queries.ListItems(ctx, pageOffset)
	if err != nil {
		return nil, nil, fmt.Errorf("GetTable issue getting item data: %w", err)
	}

	for _, v := range items {
		row := []string{
			v.Item,
			fmt.Sprintf("%05.2f", v.Price),
		}
		data = append(data, row)
	}

	headers = []string{"NAME", "PRICE"}
	return headers, data, err
}

func (i *ItemStore) GetRows(
	ctx context.Context, pageOffset int64,
) (
	[]Item, error,
) {
	return i.queries.ListItems(ctx, pageOffset)
}

func (i *ItemStore) CountRows(ctx context.Context) (
	int64, error,
) {
	return i.queries.CountItems(ctx)
}

func (i *ItemStore) Post(
	ctx context.Context, item Item,
) error {
	params := CreateItemParams{}
	params.ID = item.ID
	params.Item = item.Item
	params.Price = item.Price

	return i.queries.CreateItem(ctx, params)
}

func (i *ItemStore) Delete(
	ctx context.Context, item Item,
) error {
	return i.queries.DeleteItem(ctx, item.ID)
}

func (i *ItemStore) Parse(*huh.Form) (Item, error) { return Item{}, nil }
func (i *ItemStore) NewForm(...Item) *huh.Form     { return huh.NewForm() }
func (i *ItemStore) DeletFrom(Item) *huh.Form      { return huh.NewForm() }
