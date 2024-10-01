package ui

import (
	"fmt"

	"github.com/awesome-gocui/gocui"
)

func mdFileListView(g *gocui.Gui, maxX int, maxY int) error {
	if v, err := g.SetView("list", 0, 0, maxX/5, maxY-4, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Frame = true
		v.Title = " Markdown Files "
		v.Highlight = true
		v.SelBgColor = gocui.ColorBlue
		v.SelFgColor = gocui.ColorWhite
		v.FrameRunes = []rune{'─', '│', '╭', '╮', '╰', '╯'}

		// Set current view
		if _, err := g.SetCurrentView("list"); err != nil {
			return err
		}

		go refreshList(g)
	}

	return nil
}

func mdReaderView(g *gocui.Gui, maxX int, maxY int) error {
	if v, err := g.SetView("reader", maxX/5+1, 0, maxX-1, maxY-4, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = fmt.Sprintf(" Markdown Reader [ %d%% ]", scrollPercentage)
		v.Frame = true
		v.FrameRunes = []rune{'─', '│', '╭', '╮', '╰', '╯'}

		go refreshContent(g)
	}

	return nil
}

func statusBarView(g *gocui.Gui, maxX int, maxY int) error {
	if v, err := g.SetView("status", 0, maxY-4, maxX, maxY-2, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.BgColor = gocui.ColorBlue
		v.FgColor = gocui.ColorWhite

		go refreshStatusBar(g)
	}

	return nil
}
