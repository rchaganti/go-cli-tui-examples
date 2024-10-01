package ui

import (
	"github.com/awesome-gocui/gocui"
)

func setKeyBindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	if err := g.SetKeybinding("reader", gocui.KeyArrowDown, gocui.ModNone, scrollDown); err != nil {
		return err
	}
	if err := g.SetKeybinding("reader", gocui.KeyArrowUp, gocui.ModNone, scrollUp); err != nil {
		return err
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
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
