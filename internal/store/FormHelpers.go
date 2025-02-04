package store

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/huh"
)

func CleanInput(text string) (string, error) {
	return strings.ToLower(strings.TrimSpace(text)), nil
}

func IsFloat(s string) error {
	if _, err := strconv.ParseFloat(s, 64); err != nil {
		return errors.New("inputted price is not a float")
	}
	return nil
}

func DeleteForm(title string) *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Key("confirm").
				Affirmative("Yup").
				Negative("Nah").
				Title(fmt.Sprintf("Delete %s?", title)),
		),
	)
}
