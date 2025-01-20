package models

import (
	"fmt"
	"mathing/internal/store"

	tea "github.com/charmbracelet/bubbletea"
)

type listMenu struct {
	state    state
	store    *store.Queries
	selected int
	itemIDs  []int
	content  []string
}

func (l *listMenu) View() string {
	// body
	s := ""
	for i, string := range l.content {
		line := fmt.Sprintf("%02d. ", i+1) + string
		if i == l.selected {
			line = highlightStyle.Render(line)
		}
		s += line + "\n"
	}

	return s
}

func (l *listMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if l.selected > 0 {
				l.selected--
			}
		case "down", "j":
			if l.selected < len(l.content)-1 {
				l.selected++
			}
		}
	}
	return l, nil
}

func (l *listMenu) Init() tea.Cmd {
	return nil
}
