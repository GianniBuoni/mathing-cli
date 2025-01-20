package models

import (
	"context"
	"fmt"
	"mathing/internal/store"
)

type Receipt struct {
	tableMenu
}

func NewReceipt(s *store.Queries) (subModel, error) {
	menu := &Receipt{
		tableMenu{
			state:       listReceipt,
			store:       s,
			selected:    0,
			offset:      0,
			offsetSteps: 40,
			headers:     []string{"ITEM NAME", "ITEM PRICE", "ITEM QTY", "WHO PAYS?"},
		},
	}

	// fetched data
	if err := menu.Get(); err != nil {
		return nil, err
	}
	return menu, nil
}

func (r *Receipt) Get() error {
	ctx := context.Background()
	res, err := r.store.ListReceipt(ctx)
	if err != nil {
		return fmt.Errorf("issue getting receipt rows: %w", err)
	}

	count, err := r.store.CountReceipt(ctx)
	if err != nil {
		return fmt.Errorf("issue counting receipt rows: %w", err)
	}

	content := [][]string{}
	itemIDs := []int{}

	for _, v := range res {
		row := []string{
			v.ItemName,
			fmt.Sprintf("%05.2f", v.ItemPrice),
			fmt.Sprintf("%02d", v.ItemQty.Int64),
			v.Payee,
		}

		content = append(content, row)
		itemIDs = append(itemIDs, int(v.ID))
	}

	r.content = content
	r.itemIDs = itemIDs
	r.itemCount = int(count)
	return nil
}

func (r *Receipt) Upsert() error {
	return nil
}

func (r *Receipt) Delete() error {
	return nil
}

func (r *Receipt) NextState() state {
	return mainMenu
}
