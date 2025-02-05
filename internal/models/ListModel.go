package models

import (
	"mathing/internal/interfaces"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type ListState uint

const (
	table ListState = iota
	form
)

type ListAction uint

const (
	none ListAction = iota
	create
	remove
	edit
)

type ListModel[T any] struct {
	state       ListState
	action      ListAction
	table       *TableData
	form        *huh.Form
	actionFuncs map[ListAction]func() tea.Cmd
	store       interfaces.Store[T]
	rows        []T
}

func (lm *ListModel[T]) Init() tea.Cmd {
	return tea.Batch(lm.table.Init())
}

func (lm *ListModel[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	// SHARED UPDATE
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			lm.state = table
		}
	}
	// STATE UPDATES
	switch lm.state {
	case form:
		form, cmd := lm.form.Update(msg)
		if f, ok := form.(*huh.Form); ok {
			lm.form = f
			if lm.form.State == huh.StateCompleted &&
				lm.form.GetBool("confirm") == true {
				cmd = lm.actionFuncs[lm.action]()
			}
			cmds = append(cmds, cmd)
		}
	default:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "d":
				cmds = append(cmds, lm.DeleteInit(lm.GetCurrent()))
			case "a":
				cmds = append(cmds, lm.PostInit(create))
			case "e":
				cmds = append(cmds, lm.PostInit(edit, lm.GetCurrent()))
			}
		}
		table, cmd := lm.table.Update(msg)
		if t, ok := table.(*TableData); ok {
			lm.table = t
			cmds = append(cmds, cmd)
			cmd = lm.Refetch()
			cmds = append(cmds, cmd)
		}
	}
	return lm, tea.Batch(cmds...)
}

func (lm *ListModel[T]) View() string {
	switch lm.state {
	case form:
		if lm.form.State == huh.StateCompleted {
			lm.state = table
			lm.Refetch()
			return lm.table.View()
		}
		return lm.form.View()
	default:
		return lm.table.View()
	}
}

func (lm *ListModel[T]) RegisterAction(la ListAction, f func() tea.Cmd) {
	lm.actionFuncs[la] = f
}
