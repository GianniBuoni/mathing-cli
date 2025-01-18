package main

import (
	"log"
	"os"

	"mathing/internal/models"
	"mathing/internal/store"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	s, err := store.NewStore()
	if err != nil {
		log.Fatalf("issue initializing store: %v\n", err)
	}

	config, err := models.NewConfig(s)
	if err != nil {
		log.Fatalf("issue initializing model config: %v\n", err)
	}

	p := tea.NewProgram(&config)
	if _, err := p.Run(); err != nil {
		log.Fatalf("could not load program %v\n", err)
		os.Exit(1)
	}
}
