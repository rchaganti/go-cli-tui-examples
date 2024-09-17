package ui

import "github.com/charmbracelet/lipgloss"

const (
	width  = 25
	height = 10
)

var (
	listStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Height(height+10).
			Width(width+2).
			BorderForeground(lipgloss.Color("19")).
			Padding(1, 2)

	helpKeysStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF00FF"))
	helpDescStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#AAAFFF"))

	selectedTitle = lipgloss.NewStyle().
			Border(lipgloss.ThickBorder(), false, false, false, true).
			BorderForeground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
			Padding(0, 0, 0, 1)

	selectedDesc = selectedTitle.
			Foreground(lipgloss.AdaptiveColor{Light: "#000", Dark: "#FFF"})

	normalTitle = lipgloss.NewStyle().
			Border(lipgloss.ThickBorder(), false, false, false, true).
			BorderForeground(lipgloss.AdaptiveColor{Light: "#FF00FF", Dark: "#EE6FF8"}).
			Padding(0, 0, 0, 1)

	normalDesc = normalTitle.
			Foreground(lipgloss.AdaptiveColor{Light: "#FF00FF", Dark: "#AAAFFF"})

	dimmedTitle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#A49FA5", Dark: "#777777"}).
			Padding(0, 0, 0, 2)

	dimmedDesc = dimmedTitle.
			Foreground(lipgloss.AdaptiveColor{Light: "#C2B8C2", Dark: "#4D4D4D"})
)
