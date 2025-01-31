package main

import (
	"log"
	"mathing/internal/commands"
	"mathing/internal/store"
	"os"
)

func main() {
	state := &commands.State{}
	store, err := store.NewStore()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	state.Store = store

	commandList := commands.NewRegistry()
	commandList.Load()

	input := os.Args
	if len(input) < 2 {
		log.Printf("❌: expecting command name and command argument.")
		os.Exit(1)
	}

	command := commands.Command{
		Name: input[1],
		Args: input[2:],
	}

	if err := commandList.Run(state, command); err != nil {
		log.Fatalf("❌: issue running command. %v", err)
	}
}
