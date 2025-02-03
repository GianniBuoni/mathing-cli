package models

import (
	"fmt"
	"mathing/internal/lib"

	tea "github.com/charmbracelet/bubbletea"
)

var confirm string

func (i *ItemModel) Delete() tea.Cmd {
	title := fmt.Sprintf("Delete %s?", i.CurrentItem().Item)
  i.form = lib.NewDeleteForm(title)
	i.state = form
  return i.form.Init()
}

func (i *ItemModel) Create() {}

func (i *ItemModel) Edit() {}
