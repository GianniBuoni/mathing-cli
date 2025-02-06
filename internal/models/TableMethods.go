package models

import (
	"math"

	tea "github.com/charmbracelet/bubbletea"
)

// UPDATE
func (table *TableData) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// navigation
		case "j", "down":
			table.SelectNext()
		case "k", "up":
			table.SelectPrev()
		case "l", "right":
			table.PageNext()
		case "h", "left":
			table.PagePrev()
		}
	}
	return table, nil
}

// GETTERS
func (t *TableData) CurrentPage() int64 {
	return (t.pageOffset / 20) + 1
}

func (t *TableData) PageCount() int64 {
	return int64(math.Ceil(float64(t.itemCount) / 20))
}

// SETTERS
func (t *TableData) SelectNext() {
	if t.selected < len(t.data)-1 {
		t.selected++
	} else {
		t.selected = 0
	}
}
func (t *TableData) SelectPrev() {
	if t.selected > 0 {
		t.selected--
	} else {
		t.selected = len(t.data) - 1
	}
}
func (t *TableData) PageNext() {
	if t.CurrentPage() < t.PageCount() {
		t.pageOffset += 20
	} else {
		t.pageOffset = 0
	}
}
func (t *TableData) PagePrev() {
	if t.CurrentPage() > 1 {
		t.pageOffset -= 20
	} else {
		t.pageOffset = (t.PageCount() - 1) * 20
	}
}
