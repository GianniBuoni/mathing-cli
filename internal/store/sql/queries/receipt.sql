-- name: CreateReceipt :exec
INSERT INTO receipt (
  item_id, item_qty
) VALUES ( ?, ? )
  ON CONFLICT (item_id) DO UPDATE
  SET item_qty = excluded.item_qty, user_id = excluded.user_id;

-- name: CreateReceiptUsers :exec
INSERT INTO receipts_users (
  receipt_id, user_id
) VALUES ( ?, ? );

-- name: CountReceipt :one
SELECT count(*) FROM receipt;

-- name: CalcReceiptTotal :one
SELECT sum(i.price * r.item_qty) as calced_total
FROM receipt r
INNER JOIN items i
ON r.item_id = i.id;

-- name: DeleteReceipt :exec
DELETE FROM receipt
WHERE id = ?;

-- name: ListReceipt :many
SELECT
  r.id, r.item_id, i.item as item_name, i.price as item_price, r.item_qty,
  u.id as payee_id, u.name as payee
FROM receipt r
INNER JOIN items i
ON r.item_id = i.id
INNER JOIN receipts_users ru
ON r.id = ru.receipt_id
INNER JOIN users u
ON ru.user_id = u.id;
