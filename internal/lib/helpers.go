package lib

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/charmbracelet/huh"
)

type LoopOpts struct {
	Repl bool
}

func IsFloat(s string) error {
	if _, err := strconv.ParseFloat(s, 64); err != nil {
		return errors.New("inputted price is not a float")
	}
	return nil
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
		Run()
	return affirm
}

func ListSelect() (list string) {
	huh.NewSelect[string]().
		Title("Select a table to view").
		Options(
			huh.NewOption("Items", "items"),
			huh.NewOption("Users", "users"),
		).
		Value(&list).
		Run()

	return list
}
