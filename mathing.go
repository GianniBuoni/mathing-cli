package main

import (
	"fmt"
	"log"
	"mathing/internal/commands"
	"mathing/internal/store"
	"os"
)

func main() {
	state := commands.NewState()
	store, err := store.NewStore()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	state.Store = store
	state.CommandList.Load()

	input := os.Args
	command := commands.Command{}
	if len(input) < 2 {
		command = commands.Command{
			Name: "list",
		}
	} else {
		command = commands.Command{
			Name: input[1],
			Args: input[2:],
		}
	}

	fmt.Println()
	if err := state.CommandList.Run(state, command); err != nil {
		log.Fatalf("❌: issue running command. %v", err)
	}
}
