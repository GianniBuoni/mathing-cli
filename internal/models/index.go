package models

import (
	"mathing/internal/store"

	tea "github.com/charmbracelet/bubbletea"
)

type state uint

const (
	mainMenu state = iota
)

type config struct {
	store        *store.Queries
	allModels    map[state]tea.Model
	currentModel state
}


type subModelInfo struct {
	title       string
	description string
	init        func() tea.Model
}

func getIndex() map[state]subModelInfo {
	return map[state]subModelInfo{
		mainMenu: {
			title:       "Main Menu",
			description: "Select an action to perform",
			init:        NewMainMenu,
		},
	}
}
