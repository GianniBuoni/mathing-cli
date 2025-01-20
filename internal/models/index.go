package models

import (
	"mathing/internal/store"
)

type state uint

const (
	mainMenu state = iota
	listReceipt
	listItems
)

type cmd = uint

const (
	get cmd = iota
)

type subModelInfo struct {
	title       string
	description string
	init        func(*store.Queries) (subModel, error)
}

func getIndex() map[state]subModelInfo {
	return map[state]subModelInfo{
		mainMenu: {
			title:       "Main Menu",
			description: "Select an action to perform.",
			init:        NewMainMenu,
		},
		listReceipt: {
			title:       "List Receipt",
			description: "View, edit, or calc the current receipt",
			init:        NewReceipt,
		},
		listItems: {
			title:       "List Items",
			description: "View or edit items in the database.",
			init:        NewItemList,
		},
	}
}
