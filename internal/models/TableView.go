package models

import (
	"fmt"
	"mathing/internal/lib"
)

func (t *TableData) View() string {
	s := ""
	s += fmt.Sprint(lib.NewTable(t.headers, t.data, lib.WithSelection(t.selected)))
	s += fmt.Sprintln()
	s += fmt.Sprintf("Page %d of %d\n", t.CurrentPage(), t.PageCount())
	s += fmt.Sprintf("Total items: %d\n", t.itemCount)
	s += fmt.Sprintln()
	s += fmt.Sprintln(lib.HintStyle.Render("(h) - Previous Page, (l) - Next Page, (a) - Add Item, (d) - Delete Item"))
	s += fmt.Sprintln(lib.HintStyle.Render("(ctrl+c) - Quit"))
	s += fmt.Sprintln()

	return s
}
