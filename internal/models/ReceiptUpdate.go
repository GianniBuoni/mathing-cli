package models

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func (r *ReceiptModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	m, cmd := r.ListModel.Update(msg)
	if mm, ok := m.(ListModel); ok {
		r.ListModel = mm
		cmds = append(cmds, cmd)
	}

	// STATE UPDATE
	switch r.state {
	case form:
		form, cmd := r.form.Update(msg)
		if f, ok := form.(*huh.Form); ok {
			r.form = f
			if r.form.State == huh.StateCompleted &&
				r.form.GetBool("confirm") == true {
				cmd = r.actionFuncs[r.action]()
				cmds = append(cmds, cmd)
			}
			cmds = append(cmds, cmd)
		}
	default:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
      // crud acttions go here
			}
		}
		t, cmd := r.table.Update(msg)
		if tt, ok := t.(*TableData); ok {
			r.table = tt
			cmds = append(cmds, cmd)
			cmd = r.Refetch()
			cmds = append(cmds, cmd)
		}
	}
	return r, tea.Batch(cmds...)
}
