package commands

import (
	"context"
	"fmt"
)

var reset CommandData = CommandData{
	Name:        "reset",
	Description: "resets database tables (requires arguments)",
	Handler:     HandleReset,
}

func HandleReset(s *State, cmd Command) error {
	if !(len(cmd.Args) == 1) {
		return fmt.Errorf("expecting one argument.")
	}

	fmt.Println("⚡ cleaning up database.")
	switch cmd.Args[0] {
	case "items":
		ctx := context.Background()
		if err := s.Store.ResetItems(ctx); err != nil {
			return fmt.Errorf("could not reset items table: %w", err)
		}
	}

	fmt.Println("💀 database reset!")

	return nil
}
