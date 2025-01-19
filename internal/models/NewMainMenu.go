package models

import "mathing/internal/store"

func NewMainMenu(s *store.Queries) subModel {
	content := [][]string{}
	choices := getIndex()
	choiceItems := []state{}

	for k, v := range choices {
		if k == mainMenu {
			continue
		}
		row := []string{v.title, v.description}
		content = append(content, row)
		choiceItems = append(choiceItems, k)
	}

	return &tableSelect[state]{
		selected: 0,
		headers:  []string{"MENU", "DESCRIPTION"},
		content:  content,
		items:    choiceItems,
	}
}
