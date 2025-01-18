package models

import tea "github.com/charmbracelet/bubbletea"

type state uint

const (
	mainMenu state = iota
	receipt
	newReceipt
	items
	newItem
)

type subModel struct {
	name        string
	prompt string
	init        func(*config) (tea.Model, error)
	callback    func() error
}

func getSubModels() map[state]subModel {
	return map[state]subModel{
		mainMenu: {
			name:        "main menu",
			prompt: "Choose an action to perform!",
		},
		receipt: {
			name:        "receipt",
			prompt: "Current items to be calced",
		},
	}
}
