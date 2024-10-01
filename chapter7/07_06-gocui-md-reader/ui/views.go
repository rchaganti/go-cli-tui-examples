package ui

import (
	"fmt"

	"github.com/awesome-gocui/gocui"
)

func mdReaderView(g *gocui.Gui, maxX int, maxY int) error {
	if v, err := g.SetView("reader", 0, 0, maxX-1, maxY-1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// Configure view
		v.Title = fmt.Sprintf(" Markdown Reader [ %d%% ]", scrollPercentage)
		v.Frame = true
		v.FrameRunes = []rune{'─', '│', '╭', '╮', '╰', '╯'}
		v.SelBgColor = gocui.ColorBlue

		// Set current view
		if _, err := g.SetCurrentView("reader"); err != nil {
			return err
		}

		_, err := fmt.Fprintln(v, mdContent)
		if err != nil {
			return err
		}
	}

	return nil
}
