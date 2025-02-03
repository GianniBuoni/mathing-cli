package models

import (
	"fmt"
	"mathing/internal/lib"
	"mathing/internal/store"

	tea "github.com/charmbracelet/bubbletea"
)

func (lm *ListModel) CreateInit() tea.Cmd {
	lm.form = lib.NewItemForm()
	lm.state = form
	lm.action = create
	return lm.form.Init()
}

func (lm *ListModel) EditInit(s store.Item) tea.Cmd {
	lm.form = lib.NewItemForm(lib.WithStartValue(s))
	lm.state = form
	lm.action = edit
	return lm.form.Init()
}

func (lm *ListModel) DeleteInit(s string) tea.Cmd {
	title := fmt.Sprintf("Delete %s?", s)
	lm.form = lib.NewDeleteForm(title)
	lm.state = form
	lm.action = remove
	return lm.form.Init()
}
