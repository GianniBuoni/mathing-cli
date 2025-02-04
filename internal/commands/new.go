package commands

import (
	"mathing/internal/lib"
	"mathing/internal/store"
)

var newRow CommandData = CommandData{
	Name:        "new",
	Description: "Adds new item to a table. Requires argument(s)",
	Handler:     HandleNew,
}

func HandleNew(s *State, cmd Command) error {
	var list string
	if !(len(cmd.Args) == 1) {
		list = lib.ListSelect()
	} else {
		list = cmd.Args[0]
	}

	switch list {
	case "items":
		if err := lib.NewItemLoop(
			store.NewItemStore(s.Store), lib.WithRepl(true),
		); err != nil {
			return err
		}
	default:
		return lib.NoTableError(list)
	}
	return nil
}
