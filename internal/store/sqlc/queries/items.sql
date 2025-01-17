-- name: CreateItem :exec
INSERT INTO items (
  id, item, price
) VALUES ( ?, ?, ? )
  ON CONFLICT (id) DO UPDATE
  SET item = excluded.item, price = excluded.price;

-- name: ListItems :many
SELECT * FROM items;

-- name: DeleteItem :exec
DELETE FROM items
  WHERE id = ?;
