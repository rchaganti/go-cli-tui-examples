package main

import (
	"log"

	"github.com/awesome-gocui/gocui"
)

var currentViewIndex int

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("view1", maxX/4, maxY/4, 3*maxX/4, maxY/2, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "First View"
		v.Write([]byte("Hello, Gopher! Welcome to gocui tutorial!"))
	}

	if v, err := g.SetView("view2", maxX/4, maxY/2+1, 3*maxX/4, 3*maxY/4, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Second View"
		v.Write([]byte("Welcome to the multi-view example!"))
	}

	if currentViewIndex == 0 {
		g.SetCurrentView("view1")
	}

	return nil
}

func switchFocus(g *gocui.Gui, v *gocui.View) error {
	views := g.Views()
	currentViewIndex = (currentViewIndex + 1) % len(views)
	_, err := g.SetCurrentView(views[currentViewIndex].Name())
	return err
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal, false)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.SelFgColor = gocui.ColorBlue
	g.SelFrameColor = gocui.ColorGreen

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, switchFocus); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
