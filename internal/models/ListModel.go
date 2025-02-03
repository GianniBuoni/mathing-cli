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

type ListModel struct {
	state ListState
	table *TableData
	form  *huh.Form
	store interfaces.Store
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

func (*ListModel) Refetch() {}
