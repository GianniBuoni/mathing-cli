package models

import (
	"mathing/internal/store"

	tea "github.com/charmbracelet/bubbletea"
)

type config struct {
	store        *store.Queries
	allModels    map[state]tea.Model
	currentModel state
}

func NewConfig(storeQueries *store.Queries) (config, error) {
	models := map[state]tea.Model{}

	for k, v := range getIndex() {
		models[k] = v.init(storeQueries)
	}

	c := config{
		currentModel: mainMenu,
		allModels:    models,
	}

	return c, nil
}

func (c *config) Init() tea.Cmd {
	return nil
}

func (c *config) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return c, tea.Quit
		}
	}

	current, exists := c.allModels[c.currentModel]
	if !exists {
		return c, nil
	}

	_, cmd := current.Update(msg)

	return c, cmd
}

func (c *config) View() string {
	s := c.prompt()

	current, exists := c.allModels[c.currentModel]
	if !exists {
		return s + c.footer()
	}

	s += current.View() + c.footer()

	return s
}

func (c *config) prompt() string {
	return "\n" + promptStyle.Render("MATHEMATICAL!!") + "\n\n"
}

func (c *config) footer() string {
	return "\n\n" + hintStyle.Render("(ctrl+c) - Exit")
}
