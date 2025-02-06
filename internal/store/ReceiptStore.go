package store

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/charmbracelet/huh"
)

type RecieptStore struct {
	queries *Queries
}

func NewRecieptStore(q *Queries) *RecieptStore {
	return &RecieptStore{
		queries: q,
	}
}

func (r *RecieptStore) GetTable(ctx context.Context, pageOffset int64) (
	headers []string, data [][]string, err error,
) {
	res, err := r.queries.ListReceipt(ctx, pageOffset)
	if err != nil {
		return nil, nil, fmt.Errorf("GetTable issue getting receipt data: %w", err)
	}
	for _, v := range res {
		row := []string{
			v.ItemName,
			fmt.Sprintf("%05.2f", v.ItemPrice),
			fmt.Sprintf("%1d", v.ItemQty),
			v.Payee,
		}
		data = append(data, row)
	}
	headers = []string{"ITEM", "ITEM PRICE", "QTY", "PAYEE"}
	return headers, data, nil
}

func (r *RecieptStore) GetRows(ctx context.Context, pageOffset int64) (
	[]ListReceiptRow, error,
) {
	return r.queries.ListReceipt(ctx, pageOffset)
}

func (r *RecieptStore) CountRows(ctx context.Context) (int64, error) {
	return r.queries.CountReceipt(ctx)
}

func (r *RecieptStore) Post(ctx context.Context, lrr ListReceiptRow) error {
	// parse params
  if lrr.ID == 0 {
    lrr.ID = time.Now().Unix()
  }

	rParams := CreateReceiptParams{
    ID: lrr.ID,
		ItemID:  lrr.ItemID,
		ItemQty: lrr.ItemQty,
	}
  err := r.queries.CreateReceipt(ctx, rParams)
	if err != nil {
    return fmt.Errorf("could not insert into receipt table: %w", err)
	}

	userIDs, err := PayeeIDToUserID(lrr.PayeeID)
	if err != nil {
		return err
	}
	for _, userID := range userIDs {
		ruParams := CreateReceiptUsersParams{
			ReceiptID: lrr.ID,
			UserID:    userID,
		}
		err = r.queries.CreateReceiptUsers(ctx, ruParams)
		if err != nil {
      return fmt.Errorf("could not insert into reciepts_users_table: %w", err)
		}
	}
	return nil
}

func (r *RecieptStore) Delete(ctx context.Context, llr ListReceiptRow) error {
	return r.queries.DeleteReceipt(ctx, llr.ID)
}

func (r *RecieptStore) Parse(form *huh.Form, original ...ListReceiptRow) (
	lrr ListReceiptRow, err error,
) {
  itemID, ok := form.Get("item").(int64)
  if ok {
		lrr.ItemID = itemID
  }
	qty, err := strconv.ParseInt(form.GetString("qty"), 10, 64)
	if err != nil {
		return ListReceiptRow{}, err
	}
	lrr.ItemQty = qty
  
  uid, ok := form.Get("user").([]string)
  if ok {
    lrr.PayeeID = UserIDToPayeeID(uid)
  } else {
    return ListReceiptRow{}, errors.New("issue parsing form data: could not get user data.")
  }

	for _, r := range original {
		lrr.ID = r.ID
	}

  return lrr, nil
}

func (r *RecieptStore) NewForm(...ListReceiptRow) *huh.Form {
	ctx := context.Background()
	items, _ := r.queries.ListAllItems(ctx)
	itemNames := []huh.Option[int64]{}
	for _, i := range items {
		itemNames = append(
			itemNames,
			huh.NewOption(fmt.Sprintf("%s  %05.2f", i.Item, i.Price), i.ID),
		)
	}
	users, _ := r.queries.ListUsers(ctx)
	userNames := []huh.Option[string]{}
	for _, u := range users {
		// parse whether original receipt item had them selected
		userNames = append(userNames, huh.NewOption(u.Name, fmt.Sprintf("%d", u.ID)))
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
			huh.NewInput().Title("How Many?").Validate(IsInt).Key("qty"),
			huh.NewMultiSelect[string]().
				Title("Who Pays?").
				Options(userNames...).
				Key("user"),
			huh.NewConfirm().Title("All done?").Affirmative("Yup!").Negative("I guess not").Key("confirm"),
		),
	).WithTheme(huh.ThemeDracula())
}
func (r *RecieptStore) DeletFrom(llr ListReceiptRow) *huh.Form {
	return DeleteForm(fmt.Sprintf("%s", llr.ItemName))
}
