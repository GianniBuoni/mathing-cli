package lib

import (
	"github.com/charmbracelet/huh"
)

func Confirm(a string, n string) (affirm bool) {
	huh.NewConfirm().
		Title("Confirm").
		Affirmative(a).
		Negative(n).
		Value(&affirm).
		Run()
	return affirm
}
