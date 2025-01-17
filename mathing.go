package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"mathing/internal/models"
	"mathing/internal/store"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	ctx := context.Background()
	s, err := store.NewStore(&ctx)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	err = s.Queries.CreateItem(
		ctx,
		store.CreateItemParams{
			ID:    0,
			Item:  "taquitos",
			Price: 4.49,
		},
	)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	items, err := s.Queries.ListItems(ctx)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	for _, item := range items {
		fmt.Println(item.Item)
	}

	p := tea.NewProgram(models.NewConfig())
	if _, err := p.Run(); err != nil {
		log.Fatalf("could not load program %v", err)
		os.Exit(1)
	}
}
