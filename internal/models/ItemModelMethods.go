package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (i *ItemModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	// SHARED UPDATE
	m, cmd := i.ListModel.Update(msg)
	if mm, ok := m.(ListModel); ok {
		i.ListModel = mm
		cmds = append(cmds, cmd)
	}

	// STATE UPDATE
	switch i.state {
	case form:
		form, cmd := i.form.Update(msg)
		if f, ok := form.(*huh.Form); ok {
			i.form = f
			if i.form.State == huh.StateCompleted &&
				i.form.GetBool("confirm") == true {
				cmd = i.actionFuncs[i.action]()
				cmds = append(cmds, cmd)
			}
			cmds = append(cmds, cmd)
		}
	default:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "d":
				cmds = append(cmds, i.DeleteInit(i.CurrentItem().Item))
			case "a":
				cmds = append(cmds, i.CreateInit())
			case "e":
				cmds = append(cmds, i.EditInit(i.CurrentItem()))
			}
		}
		t, cmd := i.table.Update(msg)
		if tt, ok := t.(*TableData); ok {
			i.table = tt
			cmds = append(cmds, cmd)
			cmd = i.Refetch()
			cmds = append(cmds, cmd)
		}
	}
	return i, tea.Batch(cmds...)
}
