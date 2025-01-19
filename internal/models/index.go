package models

import (
	"mathing/internal/store"

	tea "github.com/charmbracelet/bubbletea"
)

type state uint

const (
	mainMenu state = iota
	listItems
)

type subModelInfo struct {
	title       string
	description string
	init        func(*store.Queries) tea.Model
}

func getIndex() map[state]subModelInfo {
	return map[state]subModelInfo{
		mainMenu: {
			title:       "Main Menu",
			description: "Select an action to perform.",
			init:        NewMainMenu,
		},
		listItems: {
			title:       "List Items",
			description: "View or edit items in the database.",
			init:        NewListItems,
		},
	}
}
