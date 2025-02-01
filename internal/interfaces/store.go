package interfaces

import "context"

type Store interface {
	GetItemTable(context.Context, int64) ([]string, [][]string, error)
	GetUserTable(context.Context) ([]string, [][]string, error)

	// sqlc
	CountItems(context.Context) (int64, error)
	DeleteItem(context.Context, int64) error
}
