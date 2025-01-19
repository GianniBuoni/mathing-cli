-- name: CreateReceiptItem :exec
INSERT INTO receipt (
  item_id, item_qty, user_id
) VALUES ( ?, ?, ? )
  ON CONFLICT (item_id) DO UPDATE
  SET item_qty = excluded.item_qty, user_id = excluded.user_id;

-- name: ListReceipt :many
SELECT
  r.id,
  i.item as item_name,
  i.price as item_price,
  r.item_qty, u.name as payee
FROM receipt r
INNER JOIN items i
INNER JOIN users u;

-- name: ListUserItems :many
SELECT item_id FROM receipt
WHERE user_id = ?;

-- name: DeletReceiptItem :exec
DELETE FROM receipt
WHERE id = ?;
