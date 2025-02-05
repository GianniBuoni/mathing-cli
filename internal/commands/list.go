package commands

import (
	"fmt"
	"mathing/internal/models"
	"mathing/internal/store"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var list CommandData = CommandData{
	Name:        "list",
	Description: "Lists table data. Requires argument(s).",
	Handler:     HandleList,
}

func HandleList(s *State, cmd Command) error {
	tabs := []string{"RECEIPT", "ITEMS"}

	receipts, _ := models.NewRecieptModel(store.NewRecieptStore(s.Store))
	items, _ := models.NewItemModel(store.NewItemStore(s.Store))

	tabContent := []tea.Model{receipts, items}

	m := models.TabModel{Tabs: tabs, TabContent: tabContent}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	return nil
}
