package interfaces

import "context"

type State interface {
	GetItemTable(context.Context, int64) ([]string, [][]string, error)
	GetUserTable(context.Context) ([]string, [][]string, error)
  GetItemCount(context.Context) (int64, error)
}
