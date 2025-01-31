package lib

import (
	"github.com/charmbracelet/lipgloss"
)

type Color string

const (
	bg  Color = "#282936" //background dark -> light
	bg1 Color = "#3a3c4e"
	bg2 Color = "#4d4f68"
	bg3 Color = "#626483"
	fg  Color = "#e9e9f4" //foreground light -> dark
	fg1 Color = "#f1f2f8"
	fg2 Color = "#f7f7fb"

	// accent colors
	blue   Color = "#62d6e8"
	cyan   Color = "#a1efe4"
	green  Color = "#00f769"
	orange Color = "#ffb86c"
	pink   Color = "#ea51b2"
	purple Color = "#b45bcf"
	red    Color = "#ff5555"
	yellow Color = "#ebff87"
)

var (
	PromptStyle = lipgloss.NewStyle().
			Bold(true).
			Italic(true).
			Foreground(lipgloss.Color(bg)).
			Background(lipgloss.Color(orange)).
			Padding(0, 1)
	HintStyle = lipgloss.NewStyle().
			Italic(true).
			Foreground(lipgloss.Color(bg3))
	HighlightStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(pink))
	TableStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(green))
	HeaderStyle = lipgloss.NewStyle().
			Align(lipgloss.Center).
			Bold(true).
			Padding(0, 1).
			Foreground(lipgloss.Color(green))
	NormalStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(fg))
)
