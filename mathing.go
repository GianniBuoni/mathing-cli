package main

import (
	"log"
	"os"

	"mathing/internal/models"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
  p := tea.NewProgram(models.NewConfig())
  if _, err := p.Run(); err != nil {
    log.Fatalf("could not load program %v", err)
    os.Exit(1)
  }
}
