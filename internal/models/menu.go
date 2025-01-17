package models

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type menu struct {
	choices  []string
	selected int
	config   *config
	name     state
}

func (m menu) Init() tea.Cmd {
	return nil
}

func (m menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
  if m.name != m.config.state {
    n, cmd := m.config.Update(msg)
    return n, cmd
  }
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "j":
			if m.selected > 0 {
				m.selected--
			}
		case "down", "k":
			if m.selected < len(m.choices) {
				m.selected++
			}
		case "enter":
			m.config.state = receipt
		case "q":
			return m, tea.Quit
    case "esc":
      m.config.state = mainMenu
		}
	}
	return m, nil
}

func (m menu) View() string {
  if m.name != m.config.state {
    return m.config.View()
  }
  
  s := m.config.prompt()
  s += getSubModels()[m.name].description + "\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if i == m.selected {
			cursor = highlightStyle.Render(">")
      choice = highlightStyle.Render(choice)
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

  s += "\n" + hintStyle.UnsetFaint().Render(
    "(q) - Exit, (esc) - Cancel, (enter) - Confirm",
  )
	return s
}
