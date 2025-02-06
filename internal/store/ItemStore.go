package store

import (
	"context"
	"fmt"
	"strconv"
	"time"

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
	params := CreateItemParams{
		ID:    item.ID,
		Item:  item.Item,
		Price: item.Price,
	}
	return i.queries.CreateItem(ctx, params)
}

func (i *ItemStore) Delete(
	ctx context.Context, item Item,
) error {
	return i.queries.DeleteItem(ctx, item.ID)
}

func (i *ItemStore) Parse(form *huh.Form, orignals ...Item) (parsed Item, err error) {
	parsed.Item, err = CleanInput(form.GetString("item"))
	if err != nil {
		return Item{}, fmt.Errorf("issue parsing form data: %w", err)
	}
	parsed.Price, err = strconv.ParseFloat(form.GetString("price"), 64)
	if err != nil {
		return Item{}, fmt.Errorf("issue parsing from data: %w", err)
	}
	for _, original := range orignals {
		parsed.ID = original.ID
	}
	if parsed.ID == 0 {
		parsed.ID = time.Now().Unix()
	}
	return parsed, nil
}

func (i *ItemStore) NewForm(items ...Item) *huh.Form {
	defaultV := Item{}

	for _, i := range items {
		defaultV = i
	}
	return huh.NewForm(
		newItemPrompt(defaultV),
	).WithTheme(huh.ThemeDracula())
}

func (i *ItemStore) DeletFrom(item Item) *huh.Form {
	return DeleteForm(fmt.Sprintf("%s", item.Item))
}
