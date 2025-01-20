package models

import (
	"fmt"
	"mathing/internal/store"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type tableMenu struct {
	state       state
	store       *store.Queries
	selected    int
	itemCount   int
	offset      int //page offset
	offsetSteps int
	itemIDs     []int // underlying id content
	headers     []string
	content     [][]string
}

func (t *tableMenu) View() string {
	// body
	u := table.New().Border(lipgloss.NormalBorder()).BorderStyle(tableStyle).StyleFunc(
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

	s := fmt.Sprint(u)

	if t.itemCount > 0 {
		// pagination
		pageNumbers := t.itemCount/t.offsetSteps + 1
		currentPage := t.offset/t.offsetSteps + 1
		pagination := hintStyle.Margin(0, 1).
			Render(fmt.Sprintf("Page %02d of %02d", currentPage, pageNumbers))
		s += "\n" + pagination

		// item count
		countString := hintStyle.Render(fmt.Sprintf("Count: %d items", t.itemCount))
		s += countString

		// key bindings
		s += "\n\n" + hintStyle.Render("(h) - Previous Page, (l) - Next page")
	}

	return s
}

func (t *tableMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if t.selected > 0 {
				t.selected--
			}

		case "down", "j":
			if t.selected < len(t.content)-1 {
				t.selected++
			}

		case "left", "h":
			if t.offset >= t.offsetSteps {
				t.offset -= t.offsetSteps
			}
		case "right", "l":
			if (t.offset + t.offsetSteps) <= t.itemCount {
				t.offset += t.offsetSteps
			}
		}

	}
	return t, nil
}

func (t *tableMenu) Init() tea.Cmd {
	return nil
}
