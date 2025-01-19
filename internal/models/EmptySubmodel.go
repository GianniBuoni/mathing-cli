package models

import "mathing/internal/store"

// boilerplate for new submodles
type BoilerPlateMenu struct {
	tableMenu //replace with any struct that implements the tea.Model interface
}

func NewBoilerPlateMenu(s *store.Queries) subModel {
	// hard coded data members
	menu := &BoilerPlateMenu{
		tableMenu{},
	}

	// fetched data
	menu.Get()
	return menu
}

func (m *BoilerPlateMenu) Get() error {
	return nil
}

func (m *BoilerPlateMenu) Upsert() error {
	return nil
}

func (m *BoilerPlateMenu) Delete() error {
	return nil
}

func (m *BoilerPlateMenu) NextState() state {
	return mainMenu
}
