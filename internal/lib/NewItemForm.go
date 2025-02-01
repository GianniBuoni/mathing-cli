package lib

import (
	"errors"
	"fmt"
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

func NewItemForm() (store.CreateItemParams, error) {
	data := NewItemFormData{}
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("ITEM NAME?").Value(&data.Item),
			huh.NewInput().Title("ITEM PRICE?").Validate(IsFloat).Value(&data.Price),
		),
	)

	if err := form.Run(); err != nil {
		return store.CreateItemParams{}, fmt.Errorf("form error: %w", err)
	}

	parsedData := store.CreateItemParams{
		ID: time.Now().Unix(),
	}

	parsedData.Item, _ = CleanInput(data.Item)
	parsedData.Price, _ = strconv.ParseFloat(data.Price, 64)

	return parsedData, nil
}

func IsFloat(s string) error {
	if _, err := strconv.ParseFloat(s, 64); err != nil {
		return errors.New("inputted price is not a float")
	}
	return nil
}
