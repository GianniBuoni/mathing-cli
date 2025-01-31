package commands

import (
	"context"
	"fmt"
	"mathing/internal/store"
)

var seed CommandData = CommandData{
	Name:        "seed",
	Description: "Seed the database with dummy data.",
	Handler:     HandleSeed,
}

func HandleSeed(s *State, cmd Command) error {
	ctx := context.Background()

	fmt.Println("⚡ starting seeding database.")

	for _, item := range items {
		fmt.Printf("Adding: %s\n", item.Item)
		if err := s.Store.CreateItem(ctx, item); err != nil {
			return fmt.Errorf("could not seed items: %w", err)
		}
	}

	fmt.Println("🍃database seeded!")

	return nil
}

type item = store.CreateItemParams

var items []item = []item{
	{
		Item:  "banana",
		Price: 0.49,
	},
	{
		Item:  "popcorn with herbs",
		Price: 3.49,
	},
	{
		Item:  "peeled tomatoes",
		Price: 2.49,
	},
	{
		Item:  "tomato paste",
		Price: 1.99,
	},
}
