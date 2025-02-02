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
  s += fmt.Sprintf("Total items: %d\n", i.itemCount)
	s += fmt.Sprintln()
	s += fmt.Sprintln(lib.HintStyle.Render("(h) - Previous Page, (l) - Next Page, (a) - Add Item, (d) - Delete Item"))
	s += fmt.Sprintln(lib.HintStyle.Render("(ctrl+c) - Quit"))
	s += fmt.Sprintln()

	return s
}
