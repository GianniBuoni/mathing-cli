package models

import (
	"context"
	"mathing/internal/lib"

	tea "github.com/charmbracelet/bubbletea"
)

func (i *ItemsList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			if i.selected < len(i.data)-1 {
				i.selected++
			}
		case "k", "up":
			if i.selected > 0 {
				i.selected--
			}
		case "l", "right":
			if i.CurrentPage() < i.PageCount() {
				i.pageOffset += 20
				if err := i.Refetch(); err != nil {
					return i, tea.Println(err)
				}
			}
		case "h", "left":
			if i.CurrentPage() > 1 {
				i.pageOffset -= 20
				if err := i.Refetch(); err != nil {
					return i, tea.Println(err)
				}
			}

		// CRUD actions
		case "d":
			if lib.Confirm("Delete item", "Cancel") {
				if err := i.store.DeleteItem(
					context.Background(),
					i.CurrentItem().ID,
				); err != nil {
					return i, tea.Println(err)
				}

				if err := i.Refetch(); err != nil {
					return i, tea.Println(err)
				}
				return i, nil
			}

		case "a":
			for {
				newItem, err := lib.NewItemForm()
				if err != nil {
					return i, tea.Println(err)
				}

				err = i.store.CreateItem(context.Background(), newItem)
				if err != nil {
					return i, tea.Println(err)
				}

				if lib.Confirm("All done", "Add another item") {
					break
				} else {
					continue
				}

			}
			if err := i.Refetch(); err != nil {
				return i, tea.Println(err)
			}
			return i, nil

		case "ctrl+c":
			return i, tea.Quit
		}
	}
	return i, nil
}
