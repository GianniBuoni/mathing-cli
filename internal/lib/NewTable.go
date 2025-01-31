package lib

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func NewTable(headers []string, data [][]string) *table.Table {
	return table.New().
		BorderStyle(TableStyle).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return HeaderStyle
			case row%2 == 0:
				return NormalStyle.Margin(0, 1)
			default:
				return NormalStyle.Faint(true).Margin(0, 1)
			}
		}).
		Headers(headers...).
		Rows(data...)
}
