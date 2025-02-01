package models

import (
	"context"
	"math"
)


func (i *ItemsList) Refetch() error {
	ctx := context.Background()
	_, data, err := i.state.GetItemTable(ctx, i.pageOffset)
	if err != nil {
		return err
	}

	count, err := i.state.GetItemCount(ctx)
	if err != nil {
		return err
	}

	i.data = data
	i.itemCount = count
  i.selected = 0
	return nil
}

func (i *ItemsList) CurrentPage() int64 {
	return (i.pageOffset / 20) + 1
}

func (i *ItemsList) PageCount() int64 {
	return int64(math.Ceil(float64(i.itemCount)/ 20))
}
