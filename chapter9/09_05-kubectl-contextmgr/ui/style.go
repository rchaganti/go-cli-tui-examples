package ui

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

const (
	selectedFgDark  = "#FFF"
	selectedFgLight = "#000"
	selectedBgDark  = "#AAA"
	selectedBgLight = "#FFF"
	borderFg        = "#AAA"
	normalTitle     = "#FF00FF"
	normalTitleDark = "#EE6FF8"
	normalDescDark  = "#AAAFFF"
)

// style definitions
var (
	helpStyle = lipgloss.NewStyle().
			AlignHorizontal(lipgloss.Left)

	helpKeysStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(normalTitle))
	helpDescStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(normalDescDark))
)

// style functions
func newTableStyle() table.Styles {
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(borderFg)).
		BorderBottom(true).
		BorderTop(true)

	s.Selected = s.Selected.
		Foreground(lipgloss.AdaptiveColor{Light: selectedFgLight, Dark: selectedFgDark}).
		Background(lipgloss.AdaptiveColor{Light: selectedBgLight, Dark: selectedBgDark}).
		Bold(false)

	return s
}
