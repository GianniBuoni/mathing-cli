package lib

import (
	"context"
	"fmt"
	"mathing/internal/store"
)

func NewItemLoop(s *store.ItemStore, opts ...func(*LoopOpts)) error {
	config := &LoopOpts{Repl: false}
	for _, opt := range opts {
		opt(config)
	}

	for {
		ctx := context.Background()
		form := s.NewForm()
		if err := form.Run(); err != nil {
			return fmt.Errorf("form error: %w", err)
		}
		data, err := s.Parse(form)
		if err != nil {
			return err
		}

		err = s.Post(ctx, data)
		if err != nil {
			return fmt.Errorf("issue adding new item: %w", err)
		}

		if config.Repl {
			fmt.Println("⭐ NEW!")
			fmt.Printf("Item: %s, Price: %05.2f\n", data.Item, data.Price)
			fmt.Println()
		}

		if Confirm("All done!", "Add another item.") {
			break
		} else {
			continue
		}
	}
	return nil
}
