package styles

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

var (
	textStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#000000", Dark: "#FFFFFF"})

	highlightStyle = textStyle.Bold(true).
			Align(lipgloss.Center)

	rowStyle = textStyle.
			Padding(0, 1).
			Align(lipgloss.Left)

	borderStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: "#4455BB", Dark: "#11EEFF"})
)

func RenderTable(header []string, data [][]string) {
	t := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderStyle(borderStyle).
		StyleFunc(cellStyle).
		Headers(header...).
		Rows(data...)

	fmt.Println(t)
}

func cellStyle(row, col int) lipgloss.Style {
	switch {
	case row == 0 || col == 0:
		return highlightStyle
	default:
		return rowStyle
	}
}
