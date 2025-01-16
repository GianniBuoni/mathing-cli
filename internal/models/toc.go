package models

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type toc struct {
	choices  []string
	selected int
}

func (t toc) Init() tea.Cmd {
	return nil
}

func (t toc) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return t, tea.Quit
		case "up", "k":
			if t.selected > 0 {
				t.selected--
			}
		case "down", "j":
			if t.selected < len(t.choices) {
				t.selected++
			}
		}
	}
	return t, nil
}

func (t toc) View() string {
	s := "\n" + promptStyle.Render(promptMessage) + "\n\n"

	for i, choice := range t.choices {
		cursor := " "
		if i == t.selected {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

  s += "\n" + hintStyle.Render("(q) - Exit, (enter) - confirm selection")

	return s
}

func newTOC() tea.Model {
	return toc{
		choices:  []string{"Update receipt", "New receipt", "Update item", "New item"},
		selected: 0,
	}
}
