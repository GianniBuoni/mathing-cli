package models

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
)

func (lm *ListModel[T]) Refetch() tea.Cmd {
	ctx := context.Background()
	var err error
	// get count first
	lm.table.itemCount, err = lm.store.CountRows(ctx)
	if err != nil {
		return tea.Println(err)
	}

	// change page offset if needed
	if lm.table.CurrentPage() > lm.table.PageCount() {
		lm.table.pageOffset = (lm.table.PageCount() - 1) * 20
	}

	// get new table data
	_, lm.table.data, err = lm.store.GetTable(ctx, lm.table.pageOffset)
	if err != nil {
		return tea.Println(err)
	}

	// get new rows slice
	lm.rows, err = lm.store.GetRows(ctx, lm.table.pageOffset)
	if err != nil {
		return tea.Println(err)
	}

	if lm.table.selected > len(lm.table.data)-1 {
		lm.table.selected = len(lm.table.data) - 1
	}

	return nil
}

func (lm *ListModel[T]) GetCurrent() T {
	return lm.rows[lm.table.selected]
}
