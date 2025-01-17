package models

import tea "github.com/charmbracelet/bubbletea"

type state uint

const (
	mainMenu state = iota
	receipt
)

type config struct {
	state     state
	prompt    func() string
	subModels map[state]tea.Model
}

func NewConfig() config {
	c := config{
		state:     mainMenu,
		prompt:    prompt,
		subModels: map[state]tea.Model{},
	}
	for k, v := range getSubModels() {
		c.subModels[k] = v.init(&c)
	}
	return c
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
