package models

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Table interface {
	tea.Model
	SelectNext()
	SelectPrev()
	PageNext()
	PagePrev()
	Refetch() (tea.Model, tea.Cmd)
}

func UpdateGenerator() func(tea.Msg, Table) (tea.Model, tea.Cmd) {
	return func(msg tea.Msg, table Table) (tea.Model, tea.Cmd) {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			// navigation
			case "j", "down":
				table.SelectNext()
			case "k", "up":
				table.SelectPrev()
			case "l", "right":
				table.PageNext()
				table.Refetch()
			case "h", "left":
				table.PagePrev()
				table.Refetch()
			case "ctrl+c":
				return table, tea.Quit
			}
		}
		return table, nil
	}
}

func (i *ItemsList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if i.update == nil {
		fmt.Println("❗Update method not implemented: quitting program")
		return i, tea.Quit
	}
	return i.update(msg, i)
}

/*
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
			if err := lib.NewItemLoop(i.store); err != nil {
				return i, tea.Println(err)
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
*/
