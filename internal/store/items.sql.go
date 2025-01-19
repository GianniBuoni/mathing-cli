// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: items.sql

package store

import (
	"context"
)

const createItem = `-- name: CreateItem :exec
INSERT INTO items (
  id, item, price
) VALUES ( ?, ?, ? )
  ON CONFLICT (id) DO UPDATE
  SET item = excluded.item, price = excluded.price
`

type CreateItemParams struct {
	ID    int64
	Item  string
	Price float64
}

func (q *Queries) CreateItem(ctx context.Context, arg CreateItemParams) error {
	_, err := q.db.ExecContext(ctx, createItem, arg.ID, arg.Item, arg.Price)
	return err
}

const deleteItem = `-- name: DeleteItem :exec
DELETE FROM items
  WHERE id = ?
`

func (q *Queries) DeleteItem(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteItem, id)
	return err
}

const listItems = `-- name: ListItems :many
SELECT id, item, price FROM items LIMIT 20
`

func (q *Queries) ListItems(ctx context.Context) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, listItems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var i Item
		if err := rows.Scan(&i.ID, &i.Item, &i.Price); err != nil {
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
