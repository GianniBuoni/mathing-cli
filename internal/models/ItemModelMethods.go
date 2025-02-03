package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (lm *ItemModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
	)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return lm, tea.Quit
		case "d":
			cmds = append(cmds, lm.Delete())
		case "esc":
			lm.state = table
		}
	}

	switch lm.state {
	case form:
		form, cmd := lm.form.Update(msg)
		if f, ok := form.(*huh.Form); ok {
			lm.form = f
			if lm.form.State == huh.StateCompleted {
				cmds = append(cmds, tea.Println("ACTION"))
			}
			cmds = append(cmds, cmd)
		}
	default:
		t, cmd := lm.table.Update(msg)
		if tt, ok := t.(*TableData); ok {
			lm.table = tt
			cmds = append(cmds, cmd)
			lm.Refetch()

		}
	}
	return lm, tea.Batch(cmds...)
}
