// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package store

import (
	"database/sql"
)

type Item struct {
	ID    int64
	Item  string
	Price float64
}

type Receipt struct {
	ID      int64
	ItemQty sql.NullInt64
	ItemID  sql.NullInt64
	UserID  sql.NullInt64
}

type User struct {
	ID   int64
	Name string
}
