package models

import tea "github.com/charmbracelet/bubbletea"

type tableSelect struct {
	content  [][]string
	selected int
}

func NewMainMenu() tea.Model {
	content := [][]string{}
	choices := getIndex()

	for _, v := range choices {
		row := []string{v.title, v.description}
		content = append(content, row)
	}

	return &tableSelect{
		content:  content,
		selected: 0,
	}
}

func (t *tableSelect) View() string {
	return t.content[0][0]
}

func (t *tableSelect) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return t, nil
}

func (t *tableSelect) Init() tea.Cmd {
	return nil
}
