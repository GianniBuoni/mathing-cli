package models

import (
	tea "github.com/charmbracelet/bubbletea"
)

type state struct {
	name         string
	description  string
	InitCallback func() tea.Model // pass in a pointer to the store later
}

var (
	promptMessage string = promptStyle.Render("MATHEMATICAL!!")
)

func GetModel() map[string]state {
	return map[string]state{
		"toc": {
			name:         "toc",
			description:  "Default input view. Takes a command to switch to other views",
			InitCallback: newTOC,
		},
	}
}
