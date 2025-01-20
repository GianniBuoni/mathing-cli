-- name: CreateReceiptItem :exec
INSERT INTO receipt (
  item_id, item_qty, user_id
) VALUES ( ?, ?, ? )
  ON CONFLICT (item_id) DO UPDATE
  SET item_qty = excluded.item_qty, user_id = excluded.user_id;

-- name: ListReceipt :many
SELECT
  r.id,
  r.item_id,
  i.item as item_name,
  i.price as item_price,
  r.item_qty,
  r.user_id,
  u.name as payee
FROM receipt r
INNER JOIN items i
ON r.item_id = i.id
INNER JOIN users u
ON r.user_id = u.id;

-- name: CountReceipt :one
SELECT count(*) FROM receipt;

-- name: CalcReceiptTotal :one
SELECT sum(i.price * r.item_qty) as calced_total
FROM receipt r
INNER JOIN items i
ON r.item_id = i.id;

-- name: CalcUserTotal :one
SELECT r.user_id, sum(i.price * r.item_qty) as user_total
FROM receipt r
INNER JOIN items i
ON r.item_id = i.id
WHERE r.user_id = ?;

-- name: DeletReceiptItem :exec
DELETE FROM receipt
WHERE id = ?;
