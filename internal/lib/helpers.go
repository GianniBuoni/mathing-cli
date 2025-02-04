package lib

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

type LoopOpts struct {
	Repl bool
}

func WithRepl(b bool) func(*LoopOpts) {
	return func(lo *LoopOpts) { lo.Repl = b }
}

func NoTableError(s string) error {
	return fmt.Errorf("table '%s' does not exist.", s)
}

func Confirm(a string, n string) (affirm bool) {
	huh.NewConfirm().
		Title("Confirm").
		Affirmative(a).
		Negative(n).
		Value(&affirm).
		WithTheme(huh.ThemeDracula()).
		Run()
	return affirm
}

func ListSelect() (list string) {
	huh.NewSelect[string]().
		Title("Select a table to view").
		Options(
			huh.NewOption("Receipt", "receipt"),
			huh.NewOption("Items", "items"),
			huh.NewOption("Users", "users"),
		).
		Value(&list).
		WithTheme(huh.ThemeDracula()).
		Run()

	return list
}
