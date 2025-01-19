package models

import (
	"fmt"
	"mathing/internal/store"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type tableSelect[T any] struct {
	selected    int
	itemOffset  int
	itemCount   int
	headers     []string
	items       []T
	content     [][]string
	store       *store.Queries
	refetchFunc func(*store.Queries, int) ([][]string, []T, int)
}

func (t *tableSelect[T]) View() string {
	s := table.New().Border(lipgloss.NormalBorder()).BorderStyle(tableStyle).StyleFunc(
		func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return headerStyle
			case row == t.selected:
				return highlightStyle.Margin(0, 1)
			default:
				return normalStyle.Margin(0, 1)
			}
		}).Rows(t.content...).Headers(t.headers...)
	return fmt.Sprint(s) + "\n" + fmt.Sprintf("%d", t.itemOffset)
}

func (t *tableSelect[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if t.selected > 0 {
				t.selected--
			}

		case "down", "j":
			if len(t.content) != 1 && t.selected < len(t.content) {
				t.selected++
			}

		case "left", "h":
			// refetch here
			if t.itemOffset >= 20 {
				t.itemOffset -= 20
				t.content, t.items, _ = t.refetchFunc(t.store, t.itemOffset)
			}
		case "right", "l":
			if (t.itemOffset + 20) <= t.itemCount {
				t.itemOffset += 20

				t.content, t.items, _ = t.refetchFunc(
					t.store,
					t.itemOffset,
				)
			}
		}

	}
	return t, nil
}

func (t *tableSelect[T]) Init() tea.Cmd {
	return nil
}

func (t *tableSelect[T]) NextState() state {
	switch any(t.items[t.selected]).(type) {
	case state:
		return any(t.items[t.selected]).(state)
	default:
		return mainMenu
	}
}
