-- name: CreateItem :exec
INSERT INTO items (
  id, item, price
) VALUES ( ?, ?, ? )
  ON CONFLICT (id) DO UPDATE
  SET item = excluded.item, price = excluded.price;

-- name: ListItems :many
SELECT * FROM items ORDER BY item LIMIT 20 OFFSET ? ;

-- name: ListAllItems :many
SELECT * FROM items
ORDER BY item;

-- name: GetItem :one
SELECT * FROM items WHERE id = ?;

-- name: CountItems :one
SELECT count(*) FROM items;

-- name: DeleteItem :exec
DELETE FROM items
  WHERE id = ?;
