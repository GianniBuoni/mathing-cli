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

const getRowTotal = `-- name: GetRowTotal :many
SELECT 
	ru.receipt_id, ru.user_id,
  (r.item_qty*i.price / COUNT(u.id)) as total
FROM receipts_users ru
INNER JOIN receipt r ON ru.receipt_id = r.id
INNER JOIN  users u ON ru.user_id = u.id
INNER JOIN  items i on r.item_id = i.id
GROUP by ru.receipt_id, ru.user_id
`

type GetRowTotalRow struct {
	ReceiptID int64
	UserID    int64
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
		if err := rows.Scan(&i.ReceiptID, &i.UserID, &i.Total); err != nil {
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
