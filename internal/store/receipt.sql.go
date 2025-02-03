// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: receipt.sql

package store

import (
	"context"
	"database/sql"
)

const calcReceiptTotal = `-- name: CalcReceiptTotal :one
SELECT sum(i.price * r.item_qty) as calced_total
FROM receipt r
INNER JOIN items i
ON r.item_id = i.id
`

func (q *Queries) CalcReceiptTotal(ctx context.Context) (sql.NullFloat64, error) {
	row := q.db.QueryRowContext(ctx, calcReceiptTotal)
	var calced_total sql.NullFloat64
	err := row.Scan(&calced_total)
	return calced_total, err
}

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
  item_id, item_qty
) VALUES ( ?, ? )
  ON CONFLICT (item_id) DO UPDATE
  SET item_qty = excluded.item_qty, user_id = excluded.user_id
`

type CreateReceiptParams struct {
	ItemID  int64
	ItemQty int64
}

func (q *Queries) CreateReceipt(ctx context.Context, arg CreateReceiptParams) error {
	_, err := q.db.ExecContext(ctx, createReceipt, arg.ItemID, arg.ItemQty)
	return err
}

const createReceiptUsers = `-- name: CreateReceiptUsers :exec
INSERT INTO receipts_users (
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

const listReciept = `-- name: ListReciept :many
SELECT
  r.id, r.item_id, i.item as item_name, i.price as item_price, r.item_qty,
  u.id as payee_id, u.name as payee
FROM receipt r
INNER JOIN items i
ON r.item_id = i.id
INNER JOIN receipts_users ru
ON r.id = ru.receipt_id
INNER JOIN users u
ON ru.user_id = u.id
`

type ListRecieptRow struct {
	ID        int64
	ItemID    int64
	ItemName  string
	ItemPrice float64
	ItemQty   int64
	PayeeID   int64
	Payee     string
}

func (q *Queries) ListReciept(ctx context.Context) ([]ListRecieptRow, error) {
	rows, err := q.db.QueryContext(ctx, listReciept)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListRecieptRow
	for rows.Next() {
		var i ListRecieptRow
		if err := rows.Scan(
			&i.ID,
			&i.ItemID,
			&i.ItemName,
			&i.ItemPrice,
			&i.ItemQty,
			&i.PayeeID,
			&i.Payee,
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
