package commands

type Command struct {
	Name string
	Args []string
}

type CommandData struct {
	Handler     func(*State, Command) error
	Name        string
	Description string
}

type Commands struct {
	Registry map[string]CommandData
}

func NewRegistry() *Commands {
	return &Commands{
		Registry: map[string]CommandData{},
	}
}
