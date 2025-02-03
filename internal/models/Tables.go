package models

import (
	tea "github.com/charmbracelet/bubbletea"
)

type TableData struct {
	selected   int
	itemCount  int64
	pageOffset int64
	headers    []string
	data       [][]string
}

func (i *TableData) Init() tea.Cmd {
	return nil
}

func NewTableData() *TableData {
  return &TableData{
    selected: 0,
    pageOffset: 0,
  }
}
