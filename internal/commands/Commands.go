package commands

import (
	"mathing/internal/store"
)

type State struct {
	Store       *store.Queries
	CommandList *Commands
}
type Command struct {
	Name string
	Args []string
}

type CommandData struct {
	Handler     func(*State, Command) error
	Name        string
	Description string
}

type Commands struct {
	Registry map[string]CommandData
}

func NewState() *State {
	return &State{
		CommandList: NewRegistry(),
	}
}

func NewRegistry() *Commands {
	return &Commands{
		Registry: map[string]CommandData{},
	}
}
