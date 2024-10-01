package ui

import (
	"strings"

	"github.com/awesome-gocui/gocui"
)

func setKeyBindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, switchView); err != nil {
		return err
	}

	if err := g.SetKeybinding("reader", gocui.KeyArrowDown, gocui.ModNone, scrollDown); err != nil {
		return err
	}

	if err := g.SetKeybinding("reader", gocui.KeyArrowUp, gocui.ModNone, scrollUp); err != nil {
		return err
	}

	if err := g.SetKeybinding("list", gocui.KeyArrowDown, gocui.ModNone, moveListCursorDown); err != nil {
		panic(err)
	}

	if err := g.SetKeybinding("list", gocui.KeyArrowUp, gocui.ModNone, moveListCursorUp); err != nil {
		panic(err)
	}

	if err := g.SetKeybinding("list", gocui.KeyEnter, gocui.ModNone, selectFile); err != nil {
		panic(err)
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func switchView(g *gocui.Gui, v *gocui.View) error {
	if v.Name() == "list" {
		_, err := g.SetCurrentView("reader")
		if err != nil {
			return err
		}
	} else {
		_, err := g.SetCurrentView("list")
		if err != nil {
			return err
		}
	}

	go refreshStatusBar(g)
	return nil
}

func scrollUp(g *gocui.Gui, v *gocui.View) error {
	ox, oy := v.Origin()

	if oy > 0 {
		if err := v.SetOrigin(ox, oy-1); err != nil {
			return err
		}
	}

	updateScrollPercentage(g, v)

	return nil
}

func scrollDown(g *gocui.Gui, v *gocui.View) error {
	ox, oy := v.Origin()
	_, height := v.Size()

	if oy+height < len(v.ViewBufferLines()) {
		if err := v.SetOrigin(ox, oy+1); err != nil {
			return err
		}
	}

	updateScrollPercentage(g, v)

	return nil
}

func moveListCursorDown(_ *gocui.Gui, v *gocui.View) error {
	err := moveListCursor(v, 1)
	if err != nil {
		return err
	}
	return nil
}

func moveListCursorUp(_ *gocui.Gui, v *gocui.View) error {
	err := moveListCursor(v, -1)
	if err != nil {
		return err
	}
	return nil
}

func selectFile(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	l, err := v.Line(cy)
	if err != nil {
		return nil
	}

	l = strings.Trim(l, " ")
	l = strings.TrimPrefix(l, "*")
	if err := toggleFileSelection(l); err != nil {
		return err
	}

	go refreshList(g)
	go refreshContent(g)

	return nil
}
