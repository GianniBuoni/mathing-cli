package commands

import (
	"context"
	"errors"
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
	if !(len(cmd.Args) == 1) {
		return errors.New("expecting one argument.")
	}

	ctx := context.Background()
	headers := []string{}
	data := [][]string{}
	var err error

	switch cmd.Args[0] {
	case "items":
		m, err := models.NewItemsList(s)
		if err != nil {
			return err
		}

		p := tea.NewProgram(m)
		if _, err = p.Run(); err != nil {
			return err
		}
	case "users":
		headers, data, err = s.GetUserTable(ctx)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("table '%s' does not exist.", cmd.Args[0])
	}

	t := lib.NewTable(headers, data)

	fmt.Println(t)

	return nil
}
