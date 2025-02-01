package models

import (
	"context"
	"fmt"
	"math"
	"mathing/internal/interfaces"
	"mathing/internal/lib"

	tea "github.com/charmbracelet/bubbletea"
)

type ItemsList struct {
	selected   int
	itemCount  int64
	pageOffset int64
	headers    []string
	data       [][]string
	state      interfaces.State
}

func NewItemsList(s interfaces.State) (*ItemsList, error) {
	ctx := context.Background()
	headers, data, err := s.GetItemTable(ctx, 0)
	if err != nil {
		return nil, err
	}

	count, err := s.GetItemCount(ctx)
	if err != nil {
		return nil, err
	}

	return &ItemsList{
		selected:   0,
		itemCount:  count,
		pageOffset: 0,
		headers:    headers,
		data:       data,
		state:      s,
	}, nil
}

func (i *ItemsList) Init() tea.Cmd {
	return nil
}

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

func (i *ItemsList) Refetch() error {
	ctx := context.Background()
	_, data, err := i.state.GetItemTable(ctx, i.pageOffset)
	if err != nil {
		return err
	}

	count, err := i.state.GetItemCount(ctx)
	if err != nil {
		return err
	}

	i.data = data
	i.itemCount = count
  i.selected = 0
	return nil
}

func (i *ItemsList) CurrentPage() int64 {
	return (i.pageOffset / 20) + 1
}

func (i *ItemsList) PageCount() int64 {
	return int64(math.Ceil(float64(i.itemCount)/ 20))
}
