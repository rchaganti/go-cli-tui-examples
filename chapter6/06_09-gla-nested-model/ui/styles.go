package ui

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	vpWidth  = 100
	vpHeight = 20
	taWidth  = 100
	taHeight = 3
	right    = "├"
	left     = "┤"
	l        = "─"
	le       = "┐"
	re       = "└"

	slWidth  = 25
	slHeight = 20
)

var (
	faintStyle = lipgloss.NewStyle().Faint(true).BorderForeground(lipgloss.Color("#C2B8C2"))

	listFocusedStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("19")).
				Width(slWidth+2).
				Height(slHeight+10).
				Padding(1, 2)

	listFaintStyle = faintStyle.
			Border(lipgloss.RoundedBorder()).
			Width(slWidth+2).
			Height(slHeight+10).
			Padding(1, 2)

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

	blurStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#C2B8C2")).
			BorderForeground(lipgloss.Color("#C2B8C2"))

	focusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#874BFD")).
			BorderForeground(lipgloss.Color("#874BFD"))

	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return focusStyle.
			BorderStyle(b)
	}()

	titleBlurStyle = titleStyle.BorderForeground(lipgloss.Color("#C2B8C2"))

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return focusStyle.BorderStyle(b).MarginTop(1)
	}()

	infoBlurStyle = infoStyle.BorderForeground(lipgloss.Color("#C2B8C2"))

	baseStyle = focusStyle.
			Border(lipgloss.RoundedBorder()).
			Width(taWidth).
			Height(taHeight).
			AlignVertical(lipgloss.Top)

	baseBlurStyle = blurStyle.
			Border(lipgloss.RoundedBorder()).
			Width(taWidth).
			Height(taHeight).
			AlignVertical(lipgloss.Top)

	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#874BFD"))

	userMessageStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Bold(true)
	geminiMessageStyle = lipgloss.NewStyle()

	helpStyle = lipgloss.NewStyle().
			Width(vpWidth).
			AlignHorizontal(lipgloss.Left)

	helpKeysStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF00FF"))
	helpDescStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#AAAFFF"))
)
