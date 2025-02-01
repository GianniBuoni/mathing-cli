package lib

import (

	"github.com/charmbracelet/huh"
)

func AllDone() (allDone bool) {
	huh.NewConfirm().
		Title("All Done? Or add more?").
		Affirmative("All Done").
		Negative("Add another item!").
		Value(&allDone).
		Run()
	return allDone
}
