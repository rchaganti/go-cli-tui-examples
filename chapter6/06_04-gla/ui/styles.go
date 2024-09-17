package ui

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	vpWidth  = 100
	vpHeight = 20
	taWidth  = 100
	taHeight = 5
	right    = "├"
	left     = "┤"
	l        = "─"
	le       = "┐"
	re       = "└"
)

var (
	blurStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#526D82")).
			BorderForeground(lipgloss.Color("#526D82"))

	focusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#874BFD")).
			BorderForeground(lipgloss.Color("#874BFD"))

	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return focusStyle.
			BorderStyle(b)
	}()

	titleBlurStyle = titleStyle.BorderForeground(lipgloss.Color("#526D82"))

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return focusStyle.BorderStyle(b).MarginTop(1)
	}()

	infoBlurStyle = infoStyle.BorderForeground(lipgloss.Color("#526D82"))

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
)
