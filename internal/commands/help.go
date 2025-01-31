package commands

import (
	"fmt"
	"mathing/internal/lib"
	"strings"
)

var help CommandData = CommandData{
	Name:        "help",
	Description: "displays all available actions and their expected arguments",
	Handler:     HandleHelp,
}

func HandleHelp(s *State, cmd Command) error {
	helpData := [][]string{}

	for _, v := range s.CommandList.Registry {
		name := strings.ToUpper(v.Name)
		row := []string{name, v.Description}
		helpData = append(helpData, row)
	}

	t := lib.NewTable([]string{"NAME", "DESCRIPTION"}, helpData)

	fmt.Println(t)

	return nil
}
