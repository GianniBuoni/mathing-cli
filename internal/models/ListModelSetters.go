package models

import (
	"context"

	tea "github.com/charmbracelet/bubbletea"
)

func (lm *ListModel[T]) PostInit(action ListAction, values ...T) tea.Cmd {
	lm.form = lm.store.NewForm(values...)
	lm.state = form
	lm.action = action
	return lm.form.Init()
}

func (lm *ListModel[T]) DeleteInit(t T) tea.Cmd {
	lm.form = lm.store.DeletFrom(t)
	lm.state = form
	lm.action = remove
	return lm.form.Init()
}

func (lm *ListModel[T]) Post() tea.Cmd {
	res, err := lm.store.Parse(lm.form)
	if err != nil {
		return tea.Println(err)
	}

	err = lm.store.Post(context.Background(), res)
	if err != nil {
		return tea.Println(err)
	}
	if err := lm.Refetch(); err != nil {
		return tea.Println(err)
	}
	return nil
}

func (lm *ListModel[T]) Delete() tea.Cmd {
	if err := lm.store.Delete(
		context.Background(), lm.GetCurrent(),
	); err != nil {
		return tea.Println(err)
	}
	if err := lm.Refetch(); err != nil {
		return tea.Println(err)
	}
	return nil
}
