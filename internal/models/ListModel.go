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

type ListModel struct {
	state       ListState
	action      ListAction
	table       *TableData
	form        *huh.Form
	store       interfaces.Store
	actionFuncs map[ListAction]func() tea.Cmd
}

func (i *ListModel) Init() tea.Cmd {
	return tea.Batch(i.table.Init())
}

func (lm *ListModel) View() string {
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

func (i *ListModel) RegisterAction(la ListAction, f func() tea.Cmd) {
	i.actionFuncs[la] = f
}

func (*ListModel) Refetch() {}
func (*ListModel) Delete()  {}
