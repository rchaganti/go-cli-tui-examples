package ui

import "github.com/charmbracelet/lipgloss"

var (
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().
			BorderForeground(lipgloss.Color("#874BFD")).
			BorderStyle(b).
			Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.BorderStyle(b)
	}()

	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render
)
