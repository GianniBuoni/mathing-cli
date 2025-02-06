package models

import (
	"mathing/internal/lib"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TabModel struct {
	Tabs       []string
	TabContent []tea.Model
	activeTab  int
}

func (m TabModel) Init() tea.Cmd {
	return nil
}

func (m TabModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			return m, tea.Quit
		case "tab":
			if m.activeTab < len(m.Tabs)-1 {
				m.activeTab++
			} else {
				m.activeTab = 0
			}
			return m, nil
		case "shift+tab":
			if m.activeTab > 0 {
				m.activeTab--
			} else {
				m.activeTab = len(m.Tabs) - 1
			}
		}
	}
	m.TabContent[m.activeTab], cmd = m.TabContent[m.activeTab].Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

var tabStyle lipgloss.Style

func (m TabModel) View() string {
	doc := strings.Builder{}

	var renderedTabs []string

	for i, t := range m.Tabs {
		if i == m.activeTab {
			tabStyle = lib.ActiveTabStyle
		} else {
			tabStyle = lib.TabStyle
		}
		renderedTabs = append(renderedTabs, tabStyle.Render(t))
	}
	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	doc.WriteString(row)
	doc.WriteString("\n\n")
	doc.WriteString(m.TabContent[m.activeTab].View())
	return doc.String()
}
