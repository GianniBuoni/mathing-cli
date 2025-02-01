package lib

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type TableOpt struct {
	selected int
}

func NewTable(headers []string, data [][]string, opts ...func(*TableOpt)) *table.Table {
	config := &TableOpt{
		selected: -1,
	}

	for _, opt := range opts {
		opt(config)
	}

	return table.New().
		BorderStyle(TableStyle).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return HeaderStyle
			case row == config.selected:
				return HighlightStyle.Margin(0, 1)
			case row%2 == 0:
				return NormalStyle.Margin(0, 1)
			default:
				return NormalStyle.Faint(true).Margin(0, 1)
			}
		}).
		Headers(headers...).
		Rows(data...)
}

func WithSelection(i int) func(*TableOpt) {
	return func(to *TableOpt) { to.selected = i }
}
