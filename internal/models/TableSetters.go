package models

func (t *TableData) SelectNext() {
	if t.selected < len(t.data)-1 {
		t.selected++
	}
}

func (t *TableData) SelectPrev() {
	if t.selected > 0 {
		t.selected--
	}
}
func (t *TableData) PageNext() {
	if t.CurrentPage() < t.PageCount() {
		t.pageOffset += 20
	}
}
func (t *TableData) PagePrev() {
	if t.CurrentPage() > 1 {
		t.pageOffset -= 20
	}
}
