package commands

import "mathing/internal/store"

type State struct {
	Store       *store.Queries
	CommandList *Commands
}

func NewState() *State {
	return &State{
		CommandList: NewRegistry(),
	}
}
