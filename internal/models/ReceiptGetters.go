package models

import (
	"mathing/internal/store"

	tea "github.com/charmbracelet/bubbletea"
)

func (r *ReceiptModel) Refetch() tea.Cmd { return nil }
func (r *ReceiptModel) CurrentRow() store.ListReceiptRow {
	return store.ListReceiptRow{}
}
