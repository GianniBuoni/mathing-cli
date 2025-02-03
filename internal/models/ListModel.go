package models

import (
	"mathing/internal/interfaces"
	"mathing/internal/store"

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

type ItemModel struct {
	items []store.Item
  ListModel
}
