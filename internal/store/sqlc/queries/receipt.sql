-- name: CreateReceiptItem :exec
INSERT INTO receipt (
  item_qty, item_id, user_id
) VALUES ( ?, ?, ? )
  ON CONFLICT (item_id) DO UPDATE
  SET item_qty = excluded.item_qty, user_id = excluded.user_id;

-- name: ListReceipt :many
SELECT * FROM receipt;

-- name: ListUserItems :many
SELECT item_id FROM receipt
WHERE user_id = ?;

-- name: DeletReceiptItem :exec
DELETE FROM receipt
WHERE id = ?;
