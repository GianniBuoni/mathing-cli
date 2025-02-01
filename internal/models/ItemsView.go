package models

import (
	"fmt"
	"mathing/internal/lib"
)

func (i *ItemsList) View() string {
	s := ""
	s += fmt.Sprint(lib.NewTable(i.headers, i.data, lib.WithSelection(i.selected)))
	s += fmt.Sprintln()
	s += fmt.Sprintf("Page %d of %d\n", i.CurrentPage(), i.PageCount())
	s += fmt.Sprintln()
	s += fmt.Sprintln(lib.HintStyle.Render("(h) - Previous Page, (l) - Next Page"))
	s += fmt.Sprintln(lib.HintStyle.Render("(ctrl+c) - Quit"))

	return s
}

