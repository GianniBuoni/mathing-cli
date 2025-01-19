package models

import (
	"mathing/internal/store"

	tea "github.com/charmbracelet/bubbletea"
)

type subModel interface {
	Init() tea.Cmd
	View() string
	Update(tea.Msg) (tea.Model, tea.Cmd)
	NextState() state
}

type config struct {
	store         *store.Queries
	allModels     map[state]subModel
	currentModel  state
	previousModel state
}

func NewConfig(storeQueries *store.Queries) (config, error) {
	models := map[state]subModel{}

	for k, v := range getIndex() {
		models[k] = v.init(storeQueries)
	}

	c := config{
		currentModel:  mainMenu,
		previousModel: mainMenu,
		allModels:     models,
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
		case "esc":
			current := c.currentModel
			c.currentModel = c.previousModel
			c.previousModel = current
		}
	}

	current, exists := c.allModels[c.currentModel]
	if !exists {
		return c, nil
	}

	_, cmd := current.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			c.previousModel = c.currentModel
			c.currentModel = current.NextState()
		}
	}

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
	return "\n" + hintStyle.Render("(ctrl+c) - Exit, (esc) - Back")
}
