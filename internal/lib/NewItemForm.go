package lib

import (
	"context"
	"fmt"
	"mathing/internal/store"
	"strconv"
	"time"

	"github.com/charmbracelet/huh"
)

type ItemFormData struct {
	Item  string
	Price string
}

type NewUserFormData struct {
	Name string
}

func NewItemForm(opts ...func(*ItemFormData)) *huh.Form {
	defaultV := ItemFormData{}

	for _, opt := range opts {
		opt(&defaultV)
	}

	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("ITEM NAME?").Key("item").Value(&defaultV.Item),
			huh.NewInput().Title("ITEM PRICE?").Validate(IsFloat).Key("price").Value(&defaultV.Price),
			huh.NewConfirm().Affirmative("Submit").Negative("Cancel").Key("confirm"),
		).
			WithTheme(huh.ThemeDracula()),
	)
}

func NewItemParser(form *huh.Form) (store.CreateItemParams, error) {
	parsedData := store.CreateItemParams{
		ID: time.Now().Unix(),
	}

	parsedData.Item, _ = CleanInput(form.GetString("item"))
	parsedData.Price, _ = strconv.ParseFloat(form.GetString("price"), 64)

	return parsedData, nil
}

func NewItemLoop(s *store.ItemStore, opts ...func(*LoopOpts)) error {
	config := &LoopOpts{Repl: false}
	for _, opt := range opts {
		opt(config)
	}

	for {
		ctx := context.Background()
		form := NewItemForm()
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
