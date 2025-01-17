package models

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
	promptStyle = lipgloss.NewStyle().
    Bold(true).
		Italic(true).
		Foreground(lipgloss.Color(bg)).
		Background(lipgloss.Color(orange)).
		Padding(0, 1)
  hintStyle = lipgloss.NewStyle().
    Italic(true).
    Foreground(lipgloss.Color(bg3))
  highlightStyle = lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color(pink))
)
