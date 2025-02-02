-- name: CreateReceiptItem :exec
INSERT INTO receipt (
  item_id, item_qty, user_id
) VALUES ( ?, ?, ? )
  ON CONFLICT (item_id) DO UPDATE
  SET item_qty = excluded.item_qty, user_id = excluded.user_id;

-- name: CountReceipt :one
SELECT count(*) FROM receipt;

-- name: CalcReceiptTotal :one
SELECT sum(i.price * r.item_qty) as calced_total
FROM receipt r
INNER JOIN items i
ON r.item_id = i.id;

-- name: DeletReceiptItem :exec
DELETE FROM receipt
WHERE id = ?;
