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

	for _, user := range users {
		fmt.Printf("Adding %s\n", user.Name)
		if err := s.Store.CreateUser(ctx, user); err != nil {
			return fmt.Errorf("could not seed users: %w", err)
		}
	}

	for _, receipt := range receipts {
		fmt.Printf("Adding ITEM: %d, QTY: %d\n", receipt.ItemID, receipt.ItemQty)
		if err := s.Store.CreateReceipt(ctx, receipt); err != nil {
			return fmt.Errorf("could not seed receipts: %w", err)
		}
	}

	for _, ru := range rus {
		fmt.Printf("Adding RECEIPT: %d, USER: %d\n", ru.ReceiptID, ru.UserID)
		if err := s.Store.CreateReceiptUsers(ctx, ru); err != nil {
			return fmt.Errorf("could not seed receipts_users: %w", err)
		}
	}

	fmt.Println("🍃database seeded!")

	return nil
}

type item = store.CreateItemParams

var items []item = []item{
	{ID: 0, Item: "banana", Price: 0.49},
	{ID: 1, Item: "popcorn with herbs", Price: 3.49},
	{ID: 2, Item: "peeled tomatoes", Price: 2.49},
	{ID: 3, Item: "tomato paste", Price: 1.99},
}

type user = store.CreateUserParams

var users []user = []user{
	{ID: 0, Name: "jon"},
	{ID: 1, Name: "paul"},
}

type receipt = store.CreateReceiptParams

var receipts []receipt = []receipt{
	{ItemID: 1, ItemQty: 2},
	{ItemID: 3, ItemQty: 2},
}

type ru = store.CreateReceiptUsersParams

var rus []ru = []ru{
	{ReceiptID: 1, UserID: 0},
	{ReceiptID: 1, UserID: 1},
	{ReceiptID: 2, UserID: 0},
}
