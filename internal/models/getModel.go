package models

import (
	"bufio"
	"fmt"
	"mathing/internal/lib"
)

type state struct {
	name        string
	description string
	Callback    func(*bufio.Scanner) error
  Model struct{} // bubble-tea model?
}

var (
	promptMessage string = promptStyle.Render("MATHEMATICAL!!")
)

func getModel() map[string]state {
	return map[string]state{
		"default": {
			name:        "default",
			description: "Default input view. Takes a command to switch to other views",
			Callback:    prompt,
		},
	}
}

func prompt(reader *bufio.Scanner) error {
	fmt.Println()
	fmt.Printf("%s> ", promptMessage)

	// init input reading
	reader.Scan()
	err := reader.Err()
	if err != nil {
		return fmt.Errorf("issue reading stdin: %v", err)
	}

	inputs, err := lib.CleanInput(reader.Text())
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Printf("command: %s\n", inputs[0])
	return nil
}
