package store

import (
	"context"
	"fmt"
	"slices"

	"github.com/charmbracelet/huh"
)

func (r *RecieptStore) NewForm(original ...ListReceiptRow) *huh.Form {
  defaultValues := ListReceiptRow{}
  for _, lrr := range original {
    defaultValues = lrr
  }

	// set default qty value
	var defaultQty string
	if defaultValues.ItemQty != 0 {
		defaultQty = fmt.Sprintf("%d", defaultValues.ItemQty)
	}

	// get item choices
	ctx := context.Background()
	items, _ := r.queries.ListAllItems(ctx)
	itemNames := []huh.Option[int64]{}
	for _, i := range items {
		option := huh.NewOption(fmt.Sprintf("%s  %05.2f", i.Item, i.Price), i.ID)

		// check if item matches oirignal
		if i.ID == defaultValues.ItemID {
			option = option.Selected(true)
		}
		itemNames = append(itemNames, option)
	}
  
  // get user choices
	users, _ := r.queries.ListUsers(ctx)
	userNames := []huh.Option[string]{}
  defaultUsers, _ := PayeeIDToUserID(defaultValues.PayeeID)

	for _, u := range users {
    option := huh.NewOption(u.Name, fmt.Sprintf("%d", u.ID))
		// parse whether original receipt item had them selected
    if slices.Contains(defaultUsers, u.ID) {
      option = option.Selected(true)
    }
		userNames = append(userNames, option)
	}
	return huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int64]().
				Title("Add item").
				Options(itemNames...).Key("item"),
			huh.NewConfirm().
				Title("Continue?").
				Affirmative("Yup!").
				Negative("Lemme do something else"),
		),
		huh.NewGroup(
			huh.NewInput().Title("How Many?").Validate(IsInt).Key("qty").Value(&defaultQty),
			huh.NewMultiSelect[string]().
				Title("Who Pays?").
				Options(userNames...).
				Key("user"),
			huh.NewConfirm().Title("All done?").Affirmative("Yup!").Negative("I guess not").Key("confirm"),
		),
	).WithTheme(huh.ThemeDracula())
}

