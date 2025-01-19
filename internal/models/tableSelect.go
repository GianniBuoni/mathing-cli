package models

import (
	"context"
	"fmt"
	"mathing/internal/store"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type tableSelect[T any] struct {
	selected int
	headers  []string
	content  [][]string
	items    []T
}

func NewMainMenu(s *store.Queries) tea.Model {
	content := [][]string{}
	choices := getIndex()
	choiceItems := []state{}

	for k, v := range choices {
		row := []string{v.title, v.description}
		content = append(content, row)
		choiceItems = append(choiceItems, k)
	}

	return &tableSelect[state]{
		selected: 0,
		headers:  []string{"MENU", "DESCRIPTION"},
		content:  content,
		items:    choiceItems,
	}
}

func NewListItems(s *store.Queries) tea.Model {
	content := [][]string{}
	items := []int{}

	choices, err := s.ListItems(context.Background())
	if err != nil {
		return nil
	}

	for _, v := range choices {
		row := []string{v.Item, fmt.Sprintf("%.2f", v.Price)}
		content = append(content, row)
		items = append(items, int(v.ID))
	}

	return &tableSelect[int]{
		selected: 0,
		headers:  []string{"ITEM NAME", "PRICE"},
		content:  content,
		items:    items,
	}
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
	return fmt.Sprint(s)
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
			if t.selected < len(t.content) {
				t.selected++
			}
		}
	}
	return t, nil
}

func (t *tableSelect[T]) Init() tea.Cmd {
	return nil
}
