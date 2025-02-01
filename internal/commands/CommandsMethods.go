package commands

import "fmt"

func (c *Commands) Register(data CommandData) {
	c.Registry[data.Name] = data
}

func (c *Commands) Run(s *State, cmd Command) error {
	if _, ok := c.Registry[cmd.Name]; !ok {
		return fmt.Errorf("❓: cmd '%s' not found.", cmd.Name)
	}

	if err := c.Registry[cmd.Name].Handler(s, cmd); err != nil {
		return fmt.Errorf("❌: could not run '%s': %w", cmd.Name, err)
	}

	return nil
}

func (c *Commands) Load() {
	commands := []CommandData{
		seed, reset, help, list, newRow,
	}

	for _, command := range commands {
		c.Register(command)
	}
}
