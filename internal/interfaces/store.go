package interfaces

import (
	"context"

	"github.com/charmbracelet/huh"
)

type Store[T any] interface {
	GetTable(context.Context, int64) ([]string, [][]string, error)
	GetRows(context.Context, int64) ([]T, error)
	CountRows(context.Context) (int64, error)
	Post(context.Context, T) error
	Delete(context.Context, T) error

	// forms
	Parse(*huh.Form, ...T) (T, error)
	NewForm(...T) *huh.Form
	DeletFrom(T) *huh.Form
}
