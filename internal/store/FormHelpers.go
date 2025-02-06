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

func IsInt(s string) error {
	if _, err := strconv.ParseInt(s, 10, 64); err != nil {
		return errors.New("inputted price is not a integer")
	}
	return nil
}

func IsFloat(s string) error {
	if _, err := strconv.ParseFloat(s, 64); err != nil {
		return errors.New("inputted price is not a float")
	}
	return nil
}

func PayeeIDToUserID(pid string) (uids []int64, err error) {
	if pid == "" {
		return nil, errors.New("issue parsing payee ids: no users were assigned.")
	}
	pids := strings.Split(pid, ",")
	for _, pid := range pids {
		uid, err := strconv.ParseInt(pid, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("issue parsing payee ids: %w", err)
		}
		uids = append(uids, uid)
	}
	return uids, nil
}

func UserIDToPayeeID(uid []string) string {
  return strings.Join(uid, ",")
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

func newItemPrompt(i Item) *huh.Group {
	var price string
	if i.Price == 0 {
		price = ""
	} else {
		price = fmt.Sprintf("%05.2f", i.Price)
	}
	return huh.NewGroup(
		huh.NewInput().Title("ITEM NAME?").Key("item").Value(&i.Item),
		huh.NewInput().Title("ITEM PRICE?").Validate(IsFloat).Key("price").Value(&price),
		huh.NewConfirm().Affirmative("Submit").Negative("Cancel").Key("confirm"),
	)
}
