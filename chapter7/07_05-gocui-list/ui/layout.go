package ui

import "github.com/awesome-gocui/gocui"

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if err := namespacesView(g, maxX, maxY); err != nil {
		return err
	}
	return nil
}