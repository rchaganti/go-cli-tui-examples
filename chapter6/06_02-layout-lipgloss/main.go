package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

const (
	width  = 80
	height = 35
)

var (
	baseStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Align(lipgloss.Center, lipgloss.Center).
			Bold(true).
			Foreground(lipgloss.AdaptiveColor{Light: "#0000ff", Dark: "#FFFFA9"})

	summaryStyle = baseStyle.Height(percent(height, 20)).
			Width(percent(width, 60)).
			MarginLeft(5).
			SetString("This is the summary box.")

	infoStyle = baseStyle.Height(percent(height, 20)).
			Width(percent(width, 40))

	cArea1 = baseStyle.Height(percent(height, 48)).
		Width(percent(width, 40)).
		MarginLeft(5)

	cArea2 = baseStyle.Height(percent(height, 20)).
		Width(percent(width, 60))

	cArea3 = baseStyle.Height(percent(height, 20)).
		Width(percent(width, 60))
)

func percent(v, p int) int {
	return (v * p) / 100
}

func toVertical(s string) string {
	var result string
	for _, char := range s {
		result += string(char) + "\n"
	}
	return result
}

func main() {
	str1 := lipgloss.JoinHorizontal(lipgloss.Top,
		summaryStyle.Render(),
		infoStyle.Render("Information Box"),
	)

	str2 := lipgloss.JoinHorizontal(lipgloss.Top,
		cArea1.Transform(toVertical).Render("Content Area 1"),
		lipgloss.JoinVertical(
			lipgloss.Top,
			cArea2.Render("Content Area 2"),
			cArea3.Render("Content Area 3"),
		),
	)

	str3 := lipgloss.JoinVertical(lipgloss.Top, str1, str2)

	fmt.Println(str3)
}
