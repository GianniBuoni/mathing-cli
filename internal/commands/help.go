package commands

import (
	"fmt"
	"mathing/internal/lib"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
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

	t := table.New().
		BorderStyle(lib.TableStyle).
		Headers("NAME", "DESCRIPTION").
		Rows(helpData...).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return lib.HeaderStyle
			default:
				return lib.NormalStyle.Margin(0, 1)
			}
		},
		)
	fmt.Println(t)

	return nil
}
