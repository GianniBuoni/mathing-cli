package commands

import (
	"context"
	"errors"
	"fmt"
	"mathing/internal/lib"
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
		headers, data, err = s.GetItemTable(ctx)
		if err != nil {
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
