package commands

import (
	"context"
	"fmt"
	"mathing/internal/lib"
)

var reset CommandData = CommandData{
	Name:        "reset",
	Description: "Resets database tables. Requires argument(s).",
	Handler:     HandleReset,
}

func HandleReset(s *State, cmd Command) error {
	var list string
	if !(len(cmd.Args) == 1) {
		list = lib.ListSelect()
	} else {
		list = cmd.Args[0]
	}

	fmt.Println("⚡ cleaning up database.")
	ctx := context.Background()
	switch list {
	case "items":
		if err := s.Store.ResetItems(ctx); err != nil {
			return fmt.Errorf("could not reset items table: %w", err)
		}
	case "users":
		if err := s.Store.ResetUsers(ctx); err != nil {
			return fmt.Errorf("could not reset users table: %w", err)
		}
	default:
		return lib.NoTableError(list)
	}

	fmt.Println("💀 database reset!")

	return nil
}
