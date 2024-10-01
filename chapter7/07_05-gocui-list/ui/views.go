package ui

import (
	"errors"
	"fmt"
	"strings"

	"github.com/awesome-gocui/gocui"
)

func namespacesView(g *gocui.Gui, maxX int, maxY int) error {
	x1 := maxX / 4
	y1 := maxY / 4
	x2 := maxX / 2
	y2 := maxY - 15

	// Main view
	if v, err := g.SetView("namespaces", x1, y1, x2, y2, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		// Configure view
		v.Title = " Namespaces "
		v.Frame = true
		v.FrameRunes = []rune{'─', '│', '╭', '╮', '╰', '╯'}
		v.Highlight = true
		v.SelBgColor = gocui.ColorBlue

		// Display list
		go viewNamespacesRefreshList(g)

		// Set current view
		if _, err := g.SetCurrentView("namespaces"); err != nil {
			return err
		}
	}

	return nil
}

func renameView(g *gocui.Gui, v *gocui.View) error {
	var l string
	var err error

	nv, err := g.View("namespaces")
	if err != nil {
		return err
	}

	x1, y1, x2, y2 := nv.Dimensions()

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		return nil
	} else {
		l = strings.Trim(l, " ")
		l = strings.TrimPrefix(l, "*")
	}

	if v, err := g.SetView("rename", x1-10, y1+5, x2+10, y2-3, 0); err != nil {
		if !errors.Is(err, gocui.ErrUnknownView) {
			return err
		}

		fmt.Fprintln(v, l)
		v.SetCursor(len(l), 0)

		v.Title = " Rename "
		v.Editable = true
		g.Cursor = true

		if _, err := g.SetCurrentView("rename"); err != nil {
			return err
		}
	}
	return nil
}
