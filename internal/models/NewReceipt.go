package models

import (
	"context"
	"fmt"
	"mathing/internal/store"
)

func NewReceipt(s *store.Queries) subModel {
	content, items, count := fetchReceipt(s)

	return &tableSelect[int]{
		selected:  0,
		itemCount: count,
		headers:   []string{"ITEM NAME", "ITEM PRICE", "ITEM QTY", "WHO'S PAYING"},
		items:     items,
		content:   content,
	}
}

func fetchReceipt(s *store.Queries) (
	content [][]string,
	items []int,
	count int,
) {

	ctx := context.Background()
	choices, _ := s.ListReceipt(ctx)
	count = len(choices)

	for _, v := range choices {
		row := []string{
			v.ItemName,
			fmt.Sprintf("%05.2f", v.ItemPrice),
			fmt.Sprintf("%02d", v.ItemQty.Int64),
			v.Payee,
		}
		content = append(content, row)
		items = append(items, int(v.ID))
	}
	return content, items, count
}
