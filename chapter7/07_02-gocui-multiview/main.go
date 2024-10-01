package main

import (
	"log"

	"github.com/awesome-gocui/gocui"
)

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

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal, false)
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Fatal(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Fatal(err)
	}
}
