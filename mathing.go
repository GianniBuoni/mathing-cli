package main

import (
	"log"
	"mathing/internal/models"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
  m := models.GetModel()["toc"].InitCallback()
  p := tea.NewProgram(m)
  if _, err := p.Run(); err != nil {
    log.Fatalf("could not load program %v", err)
    os.Exit(1)
  }
}
