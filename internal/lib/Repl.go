package lib

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var (
	prompt string = promptStyle.Render("MATHEMATICAL!!")
)

func Repl() {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println()
		fmt.Printf("%s> ", prompt)

		// init input reading
		reader.Scan()
		err := reader.Err()
		if err != nil {
			log.Fatalf("issue reading stdin: %v", err)
		}

		inputs, err := cleanInput(reader.Text())
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			continue
		}

		fmt.Println()
		fmt.Printf("command: %s\n", inputs[0])
		continue
	}
}
