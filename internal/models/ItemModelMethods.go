package models

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (i *ListModel) Init() tea.Cmd {
  return tea.Batch(i.table.Init())
}
func (lm *ItemModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return lm, tea.Quit
		}
	}

	switch lm.state {
	case table:
		_, cmd = lm.table.Update(msg)
		cmds = append(cmds, cmd)
    lm.Refetch()
	}
	return lm, nil
}
func (lm *ListModel) View() string {
	s := lm.table.View()
	return s
}
