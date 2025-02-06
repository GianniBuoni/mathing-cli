package store

import (
	"context"
	"errors"
	"fmt"
	"strconv"

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

func (r *RecieptStore) DeletFrom(llr ListReceiptRow) *huh.Form {
	return DeleteForm(fmt.Sprintf("%s", llr.ItemName))
}
