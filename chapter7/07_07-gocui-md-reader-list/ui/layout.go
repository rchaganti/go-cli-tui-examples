package ui

import "github.com/awesome-gocui/gocui"

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if err := mdFileListView(g, maxX, maxY); err != nil {
		return err
	}

	if err := mdReaderView(g, maxX, maxY); err != nil {
		return err
	}

	if err := statusBarView(g, maxX, maxY); err != nil {
		return err
	}

	return nil
}
