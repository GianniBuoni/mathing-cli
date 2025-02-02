package models

import (
	"context"
	"math"
	"mathing/internal/store"
)

func (i *ItemsList) CurrentPage() int64 {
	return (i.pageOffset / 20) + 1
}

func (i *ItemsList) PageCount() int64 {
	return int64(math.Ceil(float64(i.itemCount) / 20))
}

func (i *ItemsList) CurrentItem() store.Item {
	return i.items[i.selected]
}

func (i *ItemsList) Refetch() (err error) {
	ctx := context.Background()
	_, i.data, err = i.store.GetItemTable(ctx, i.pageOffset)
	if err != nil {
		return err
	}

	i.items, err = i.store.ListItems(ctx, i.pageOffset)
	if err != nil {
		return err
	}

	i.itemCount, err = i.store.CountItems(ctx)
	if err != nil {
		return err
	}

	i.selected = 0
	return nil
}
