package commands

import (
	"context"
	"errors"
	"fmt"
	"mathing/internal/lib"
)

var list CommandData = CommandData{
	Name:        "list",
	Description: "Lists table data. Requires argument(s).",
	Handler:     HandleList,
}

func HandleList(s *State, cmd Command) error {
	if !(len(cmd.Args) == 1) {
		return errors.New("expecting one argument.")
	}

	tableData := [][]string{}
	headers := []string{"NAME", "PRICE"}

	ctx := context.Background()

	items, err := s.Store.ListItems(ctx, 0)
	if err != nil {
		return fmt.Errorf("issue getting item data: %w", err)
	}

	for _, v := range items {
		price := fmt.Sprintf("%05.2f", v.Price)
		row := []string{v.Item, price}
		tableData = append(tableData, row)
	}

	t := lib.NewTable(headers, tableData)

	fmt.Println(t)

	return nil
}
