package models

import (
	"mathing/internal/store"
)

type MainMenu struct {
	tableMenu
}

func NewMainMenu(s *store.Queries) subModel {
	menu := &MainMenu{
		tableMenu{
			state:    mainMenu,
			selected: 0,
			headers:  []string{"MENU", "DESCRIPTION"},
		},
	}

	menu.Get()
	return menu
}

func (m *MainMenu) Get() error {
	res := getIndex()

	itemIds := []int{}
	content := [][]string{}

	for k, v := range res {
		if k != mainMenu {
			row := []string{v.title, v.description}
			content = append(content, row)

			itemIds = append(itemIds, int(k))

		}
	}

	// update values
	m.itemIDs = itemIds
	m.content = content

	return nil
}

func (m *MainMenu) Upsert() error {
	return nil
}

func (m *MainMenu) Delete() error {
	return nil
}

func (m *MainMenu) NextState() state {
	return state(m.itemIDs[m.selected])
}
