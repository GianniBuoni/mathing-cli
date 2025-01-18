package models

import (
	"mathing/internal/store"

	tea "github.com/charmbracelet/bubbletea"
)

type state uint

const (
	mainMenu state = iota
	receipt
)

type config struct {
	state     state
	prompt    func() string
	store     *store.Queries
	subModels map[state]tea.Model
}

func NewConfig(storeQueries *store.Queries) (config, error) {
	c := config{
		state:     mainMenu,
		prompt:    prompt,
		store:     storeQueries,
		subModels: map[state]tea.Model{},
	}
	for k, v := range getSubModels() {
		var err error
		c.subModels[k], err = v.init(&c)
		if err != nil {
			return config{}, err
		}
	}
	return c, nil
}

func (c config) Init() tea.Cmd {
	return nil
}

func (c config) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	d, cmd := c.subModels[c.state].Update(msg)
	return d, cmd
}

func (c config) View() string {
	return c.subModels[c.state].View()
}

func prompt() string {
	return "\n" + promptStyle.Render("MATHEMATICAL!!") + "\n\n"
}
