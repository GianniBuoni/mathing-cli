-- name: CreateUser :exec
INSERT INTO users (
  id, name, multiplier
) VALUES (?, ?, ? )
  ON CONFLICT (id) DO UPDATE
  SET name = excluded.name, multiplier = excluded.multiplier;

-- name: ListUsers :many
SELECT * FROM users;

-- name: DelteUser :exec
DELETE FROM users
  WHERE id = ?;
