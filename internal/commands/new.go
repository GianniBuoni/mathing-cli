package commands

import (
	"context"
	"errors"
	"fmt"
	"mathing/internal/lib"
)

var newRow CommandData = CommandData{
	Name:        "new",
	Description: "Adds new item to a table. Requires argument(s)",
	Handler:     HandleNew,
}

func HandleNew(s *State, cmd Command) error {
	if !(len(cmd.Args) == 1) {
		return errors.New("expecting one argument")
	}

	for {
		ctx := context.Background()
		data, err := lib.NewItemForm()
		if err != nil {
			return err
		}

		err = s.Store.CreateItem(ctx, data)
		if err != nil {
			return fmt.Errorf("issue adding new item: %w", err)
		}

		fmt.Println("⭐ NEW!")
		fmt.Printf("Item: %s, Price: %05.2f\n", data.Item, data.Price)
		fmt.Println()

		if lib.AllDone() {
			break
		} else {
			continue
		}
	}

	return nil
}
