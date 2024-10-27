package ui

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	docWidth  = 100
	docHeight = 30

	clWidth  = 25
	clHeight = 20

	focusedBorder   = "#04B575"
	faintBorder     = "#C2B8C2"
	selectedTitle   = "#04B575"
	normalTitle     = "#FF00FF"
	normalTitleDark = "#EE6FF8"
	normalDescDark  = "#AAAFFF"
	dimTitleLight   = "#A49FA5"
	dimTitleDark    = "#777777"
	dimDescLight    = "#C2B8C2"
	dimDescDark     = "#4D4D4D"
)

// Border runes
var (
	activeTabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}

	inactiveTabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}

	extendBorder = lipgloss.Border{
		Top:         " ",
		Bottom:      "─",
		Left:        " ",
		Right:       " ",
		TopLeft:     " ",
		TopRight:    " ",
		BottomLeft:  "─",
		BottomRight: "╮",
	}
)

// style definitions
var (
	faintStyle = lipgloss.NewStyle().Faint(true).BorderForeground(lipgloss.Color(faintBorder))

	listFocusedStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color(focusedBorder)).
				Width(clWidth+2).
				Height(clHeight+10).
				Padding(1, 2)

	listFaintStyle = faintStyle.
			Border(lipgloss.RoundedBorder()).
			Width(clWidth+2).
			Height(clHeight+10).
			Padding(1, 2)

	selectedTitleStyle = lipgloss.NewStyle().
				Border(lipgloss.ThickBorder(), false, false, false, true).
				BorderForeground(lipgloss.Color(selectedTitle)).
				Padding(0, 0, 0, 1)

	selectedDescStyle = selectedTitleStyle.
				Foreground(lipgloss.AdaptiveColor{Light: "#000", Dark: "#FFF"})

	normalTitleStyle = lipgloss.NewStyle().
				Border(lipgloss.ThickBorder(), false, false, false, true).
				BorderForeground(lipgloss.AdaptiveColor{Light: normalTitle, Dark: normalTitleDark}).
				Padding(0, 0, 0, 1)

	normalDescStyle = normalTitleStyle.
			Foreground(lipgloss.AdaptiveColor{Light: normalTitle, Dark: normalDescDark})

	dimTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.AdaptiveColor{Light: dimTitleLight, Dark: dimTitleDark}).
			Padding(0, 0, 0, 2)

	dimDescStyle = dimTitleStyle.
			Foreground(lipgloss.AdaptiveColor{Light: dimDescLight, Dark: dimDescDark})

	activeTabStyle = lipgloss.NewStyle().
			BorderStyle(activeTabBorder).
			Padding(0, 1)

	inactiveTabStyle = lipgloss.NewStyle().
				BorderStyle(inactiveTabBorder).
				Padding(0, 1)

	viewerStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder(), false, true, true, false).
			Padding(0, 1)

	helpStyle = lipgloss.NewStyle().
			Width(docWidth).
			AlignHorizontal(lipgloss.Left)

	helpKeysStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(normalTitle))
	helpDescStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(normalDescDark))
)
