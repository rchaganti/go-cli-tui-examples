package ui

import "github.com/charmbracelet/lipgloss"

var (
	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render

	listStyle = lipgloss.NewStyle().Padding(1, 2)
)
