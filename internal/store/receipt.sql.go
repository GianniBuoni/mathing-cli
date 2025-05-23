// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: receipt.sql

package store

import (
	"context"
)

const countReceipt = `-- name: CountReceipt :one
SELECT count(*) FROM receipt
`

func (q *Queries) CountReceipt(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countReceipt)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createReceipt = `-- name: CreateReceipt :exec
INSERT INTO receipt (
  id, item_id, item_qty
) VALUES ( ?, ?, ? )
  ON CONFLICT (item_id) DO UPDATE
  SET item_qty = excluded.item_qty
`

type CreateReceiptParams struct {
	ID      int64
	ItemID  int64
	ItemQty int64
}

func (q *Queries) CreateReceipt(ctx context.Context, arg CreateReceiptParams) error {
	_, err := q.db.ExecContext(ctx, createReceipt, arg.ID, arg.ItemID, arg.ItemQty)
	return err
}

const createReceiptUsers = `-- name: CreateReceiptUsers :exec
INSERT OR IGNORE INTO receipts_users (
  receipt_id, user_id
) VALUES ( ?, ? )
`

type CreateReceiptUsersParams struct {
	ReceiptID int64
	UserID    int64
}

func (q *Queries) CreateReceiptUsers(ctx context.Context, arg CreateReceiptUsersParams) error {
	_, err := q.db.ExecContext(ctx, createReceiptUsers, arg.ReceiptID, arg.UserID)
	return err
}

const deleteReceipt = `-- name: DeleteReceipt :exec
DELETE FROM receipt
WHERE id = ?
`

func (q *Queries) DeleteReceipt(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteReceipt, id)
	return err
}

const deleteRecietsUsers = `-- name: DeleteRecietsUsers :exec
DELETE FROM receipts_users
WHERE receipt_id = ?
`

func (q *Queries) DeleteRecietsUsers(ctx context.Context, receiptID int64) error {
	_, err := q.db.ExecContext(ctx, deleteRecietsUsers, receiptID)
	return err
}

const getRowTotal = `-- name: GetRowTotal :many
SELECT 
	ru.receipt_id,
  GROUP_CONCAT(ru.user_id) as payees,
  (r.item_qty*i.price / COUNT(ru.user_id)) as total
FROM receipts_users ru
INNER JOIN receipt r ON ru.receipt_id = r.id
INNER JOIN  users u ON ru.user_id = u.id
INNER JOIN  items i on r.item_id = i.id
GROUP by ru.receipt_id
`

type GetRowTotalRow struct {
	ReceiptID int64
	Payees    string
	Total     interface{}
}

func (q *Queries) GetRowTotal(ctx context.Context) ([]GetRowTotalRow, error) {
	rows, err := q.db.QueryContext(ctx, getRowTotal)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRowTotalRow
	for rows.Next() {
		var i GetRowTotalRow
		if err := rows.Scan(&i.ReceiptID, &i.Payees, &i.Total); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listReceipt = `-- name: ListReceipt :many
SELECT
  r.id , r.item_id, i.item as item_name, i.price as item_price, r.item_qty,
  GROUP_CONCAT(u.name) as payee,
  GROUP_CONCAT(u.id) as payee_id,
  COUNT(u.id) as payee_count
FROM receipts_users ru
INNER JOIN receipt r
ON ru.receipt_id= r.id
INNER JOIN users u
ON ru.user_id = u.id
INNER JOIN items i
ON r.item_id = i.id
GROUP BY r.id
LIMIT 20 OFFSET ?
`

type ListReceiptRow struct {
	ID         int64
	ItemID     int64
	ItemName   string
	ItemPrice  float64
	ItemQty    int64
	Payee      string
	PayeeID    string
	PayeeCount int64
}

func (q *Queries) ListReceipt(ctx context.Context, offset int64) ([]ListReceiptRow, error) {
	rows, err := q.db.QueryContext(ctx, listReceipt, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListReceiptRow
	for rows.Next() {
		var i ListReceiptRow
		if err := rows.Scan(
			&i.ID,
			&i.ItemID,
			&i.ItemName,
			&i.ItemPrice,
			&i.ItemQty,
			&i.Payee,
			&i.PayeeID,
			&i.PayeeCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const resetReceipt = `-- name: ResetReceipt :exec
DELETE FROM receipt
`

func (q *Queries) ResetReceipt(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, resetReceipt)
	return err
}
