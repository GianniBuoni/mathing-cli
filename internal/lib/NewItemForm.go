package lib

import (
	"context"
	"fmt"
	"mathing/internal/interfaces"
	"mathing/internal/store"
	"strconv"
	"time"

	"github.com/charmbracelet/huh"
)

type NewItemFormData struct {
	Item  string
	Price string
}

type NewUserFormData struct {
	Name string
}

func NewItemForm() *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("ITEM NAME?").Key("item"),
			huh.NewInput().Title("ITEM PRICE?").Validate(IsFloat).Key("price"),
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

func NewItemLoop(s interfaces.Store, opts ...func(*LoopOpts)) error {
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
		data, err := NewItemParser(form)
		if err != nil {
			return err
		}

		err = s.CreateItem(ctx, data)
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
