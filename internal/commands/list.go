package commands

import (
	"context"
	"errors"
	"fmt"
	"mathing/internal/lib"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
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

	t := table.New().
		BorderStyle(lib.TableStyle).
		Headers(headers...).
		Rows(tableData...).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return lib.HeaderStyle
			default:
				return lib.NormalStyle.Margin(0, 1)
			}
		})

	fmt.Println(t)

	return nil
}
