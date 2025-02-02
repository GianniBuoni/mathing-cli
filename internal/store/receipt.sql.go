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

const createReceiptItem = `-- name: CreateReceiptItem :exec
INSERT INTO receipt (
  item_id, item_qty, user_id
) VALUES ( ?, ?, ? )
  ON CONFLICT (item_id) DO UPDATE
  SET item_qty = excluded.item_qty, user_id = excluded.user_id
`

type CreateReceiptItemParams struct {
	ItemID  sql.NullInt64
	ItemQty sql.NullInt64
	UserID  sql.NullInt64
}

func (q *Queries) CreateReceiptItem(ctx context.Context, arg CreateReceiptItemParams) error {
	_, err := q.db.ExecContext(ctx, createReceiptItem, arg.ItemID, arg.ItemQty, arg.UserID)
	return err
}

const deletReceiptItem = `-- name: DeletReceiptItem :exec
DELETE FROM receipt
WHERE id = ?
`

func (q *Queries) DeletReceiptItem(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletReceiptItem, id)
	return err
}
