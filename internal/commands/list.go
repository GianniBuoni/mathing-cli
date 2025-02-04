package commands

import (
	"context"
	"fmt"
	"mathing/internal/lib"
	"mathing/internal/models"

	tea "github.com/charmbracelet/bubbletea"
)

var list CommandData = CommandData{
	Name:        "list",
	Description: "Lists table data. Requires argument(s).",
	Handler:     HandleList,
}

func HandleList(s *State, cmd Command) error {
	list := ""
	if !(len(cmd.Args) == 1) {
		list = lib.ListSelect()
	} else {
		list = cmd.Args[0]
	}

	ctx := context.Background()
	headers := []string{}
	data := [][]string{}
	var err error

	switch list {
	case "items":
		m, err := models.NewItemsList(s.Store)
		if err != nil {
			return err
		}

		p := tea.NewProgram(m)
		if _, err = p.Run(); err != nil {
			return err
		}
	case "users":
		headers, data, err = s.Store.GetUserTable(ctx)
		if err != nil {
			return err
		}
  case "receipt":
    m, err := models.NewReceiptList(s.Store)
		if err != nil {
			return err
		}

		p := tea.NewProgram(m)
		if _, err = p.Run(); err != nil {
			return err
		}
	default:
		return lib.NoTableError(list)
	}

	t := lib.NewTable(headers, data)

	fmt.Println(t)

	return nil
}
