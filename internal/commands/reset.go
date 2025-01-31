package commands

import (
	"context"
	"errors"
	"fmt"
)

var reset CommandData = CommandData{
	Name:        "reset",
	Description: "Resets database tables. Requires argument(s).",
	Handler:     HandleReset,
}

func HandleReset(s *State, cmd Command) error {
	if !(len(cmd.Args) == 1) {
		return errors.New("expecting one argument.")
	}

	fmt.Println("⚡ cleaning up database.")
	ctx := context.Background()
	switch cmd.Args[0] {
	case "items":
		if err := s.Store.ResetItems(ctx); err != nil {
			return fmt.Errorf("could not reset items table: %w", err)
		}
	case "users":
		if err := s.Store.ResetUsers(ctx); err != nil {
			return fmt.Errorf("could not reset users table: %w", err)
		}
  default:
		return fmt.Errorf("table '%s' does not exist.", cmd.Args[0])
	}

	fmt.Println("💀 database reset!")

	return nil
}
