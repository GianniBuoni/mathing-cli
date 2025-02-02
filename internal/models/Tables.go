package models

import (
	"mathing/internal/interfaces"

	tea "github.com/charmbracelet/bubbletea"
)

type TableData struct {
	selected   int
	itemCount  int64
	pageOffset int64
	store      interfaces.Store
	headers    []string
	data       [][]string
}

func (i *TableData) Init() tea.Cmd {
	return nil
}
