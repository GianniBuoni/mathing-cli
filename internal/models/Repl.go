package models

import (
	"bufio"
	"fmt"
	"os"
)



func Repl() {
	reader := bufio.NewScanner(os.Stdin)

	for {
		// display appropriate view state
    viewState := "default"
    err := getModel()[viewState].Callback(reader)
    if err != nil {
      fmt.Println()
      fmt.Println(err)
    }
		continue
	}
}
