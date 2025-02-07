package store

import (
	"context"
	"fmt"
	"time"
)

func (r *RecieptStore) Post(ctx context.Context, lrr ListReceiptRow) error {
	// parse params
	if lrr.ID == 0 {
		lrr.ID = time.Now().Unix()
	}

	rParams := CreateReceiptParams{
		ID:      lrr.ID,
		ItemID:  lrr.ItemID,
		ItemQty: lrr.ItemQty,
	}
	err := r.queries.CreateReceipt(ctx, rParams)
	if err != nil {
		return fmt.Errorf("could not insert into receipt table: %w", err)
	}

  // delete receipts_users rows if there's a count mismatch
	userIDs, err := PayeeIDToUserID(lrr.PayeeID)
	if err != nil {
		return err
	}
  if int(lrr.PayeeCount) > len(userIDs) {
    err := r.queries.DeleteRecietsUsers(ctx, lrr.ID)
    if err != nil {
      return err
    }
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
