package models

import tea "github.com/charmbracelet/bubbletea"

type subModel struct {
	name        string
	description string
	init        func(*config) (tea.Model, error)
}

func getSubModels() map[state]subModel {
	return map[state]subModel{
		mainMenu: {
			name:        "main menu",
			description: "Choose an action to perform!",
			init:        newMainMenu,
		},
		receipt: {
			name:        "receipt",
			description: "Current items to be calced",
			init:        newReceipt,
		},
	}
}
