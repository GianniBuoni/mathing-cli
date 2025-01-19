package models

import (
	"context"
	"fmt"
	"mathing/internal/store"
)

func NewListItems(s *store.Queries) subModel {
	content, items, count := FetchListItems(s, 0)

	return &tableSelect[int]{
		selected:    0,
		itemOffset:  0,
		itemCount:   count,
		headers:     []string{"ITEM NAME", "PRICE"},
		items:       items,
		content:     content,
		store:       s,
		refetchFunc: FetchListItems,
	}
}

func FetchListItems(
	s *store.Queries,
	offset int,
) (
	[][]string,
	[]int,
	int,
) {
	content := [][]string{}
	items := []int{}

	count, err := s.CountItems(context.Background())
	if err != nil {
		return nil, nil, 0
	}

	choices, err := s.ListItems(context.Background(), int64(offset))
	if err != nil {
		return nil, nil, 0
	}

	for _, v := range choices {
		row := []string{v.Item, fmt.Sprintf("%05.2f", v.Price)}
		content = append(content, row)
		items = append(items, int(v.ID))
	}

	return content, items, int(count)
}
