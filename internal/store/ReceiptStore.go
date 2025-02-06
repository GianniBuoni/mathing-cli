package store

import (
	"context"
	"fmt"

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
	rParams := CreateReceiptParams{
		ItemID:  lrr.ItemID,
		ItemQty: lrr.ItemQty,
	}
	userIDs, err := PayeeIDToUserID(lrr.PayeeID)
	if err != nil {
		return err
	}

	// insert into tables
	err = r.queries.CreateReceipt(ctx, rParams)
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
			return err
		}
	}
	return nil
}

func (r *RecieptStore) Delete(ctx context.Context, llr ListReceiptRow) error {
	return r.queries.DeleteReceipt(ctx, llr.ID)
}

func (r *RecieptStore) Parse(*huh.Form, ...ListReceiptRow) (
	ListReceiptRow, error,
) {
	return ListReceiptRow{}, nil
}

func (r *RecieptStore) NewForm(...ListReceiptRow) *huh.Form {
	items, _ := r.queries.ListItems(context.Background(), 0)
	itemNames := []string{}
	for _, i := range items {
		itemNames = append(itemNames, i.Item)
	}
	return huh.NewForm(
		huh.NewGroup(),
	)
}
func (r *RecieptStore) DeletFrom(llr ListReceiptRow) *huh.Form {
	return DeleteForm(fmt.Sprintf("%s", llr.ItemName))
}
