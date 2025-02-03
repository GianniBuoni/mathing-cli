package lib

import (
	"github.com/charmbracelet/huh"
)

func NewDeleteForm(title string) *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Key("confirm").
				Affirmative("Yup").
				Negative("Nah").
				Title(title),
		),
	)
}
