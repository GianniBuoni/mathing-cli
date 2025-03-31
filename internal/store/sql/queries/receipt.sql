-- name: CreateReceipt :exec
INSERT INTO receipt (
  id, item_id, item_qty
) VALUES ( ?, ?, ? )
  ON CONFLICT (item_id) DO UPDATE
  SET item_qty = excluded.item_qty;

-- name: CreateReceiptUsers :exec
INSERT OR IGNORE INTO receipts_users (
  receipt_id, user_id
) VALUES ( ?, ? );

-- name: CountReceipt :one
SELECT count(*) FROM receipt;

-- name: DeleteReceipt :exec
DELETE FROM receipt
WHERE id = ?;

-- name: DeleteRecietsUsers :exec
DELETE FROM receipts_users
WHERE receipt_id = ?;

-- name: ResetReceipt :exec
DELETE FROM receipt;

-- name: ListReceipt :many
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
LIMIT 20 OFFSET ?;

-- name: GetRowTotal :many
SELECT 
	ru.receipt_id,
  GROUP_CONCAT(ru.user_id) as payees,
  (r.item_qty*i.price / COUNT(ru.user_id)) as total
FROM receipts_users ru
INNER JOIN receipt r ON ru.receipt_id = r.id
INNER JOIN  users u ON ru.user_id = u.id
INNER JOIN  items i on r.item_id = i.id
GROUP by ru.receipt_id;
