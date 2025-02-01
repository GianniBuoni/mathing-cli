package models

import 	tea "github.com/charmbracelet/bubbletea"

func (i *ItemsList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			if i.selected < len(i.data)-1 {
				i.selected++
			}
		case "k", "up":
			if i.selected > 0 {
				i.selected--
			}
		case "l", "right":
			if i.CurrentPage() < i.PageCount() {
				i.pageOffset += 20
				if err := i.Refetch(); err != nil {
					return i, tea.Println(err)
				}
			}
		case "h", "left":
			if i.CurrentPage() > 1 {
				i.pageOffset -= 20
				if err := i.Refetch(); err != nil {
					return i, tea.Println(err)
				}
			}

		case "ctrl+c":
			return i, tea.Quit
		}
	}
	return i, nil
}

