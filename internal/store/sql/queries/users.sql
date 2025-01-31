-- name: CreateUser :exec
INSERT INTO users (
  id, name
) VALUES (?, ?)
  ON CONFLICT (id) DO UPDATE
  SET name = excluded.name;

-- name: ListUsers :many
SELECT * FROM users;

-- name: GetUser :one
SELECT *
  FROM users
  WHERE id = ?;

-- name: DelteUser :exec
DELETE FROM users
  WHERE id = ?;
