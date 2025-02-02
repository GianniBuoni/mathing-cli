package models

func (t *TableData) SelectNext() {
	if t.selected < len(t.data)-1 {
		t.selected++
	} else {
		t.selected = 0
	}
}

func (t *TableData) SelectPrev() {
	if t.selected > 0 {
		t.selected--
	} else {
		t.selected = len(t.data) - 1
	}
}
func (t *TableData) PageNext() {
	if t.CurrentPage() < t.PageCount() {
		t.pageOffset += 20
	} else {
		t.pageOffset = 0
	}
}
func (t *TableData) PagePrev() {
	if t.CurrentPage() > 1 {
		t.pageOffset -= 20
	} else {
		t.pageOffset = (t.PageCount() - 1) * 20
	}
}
