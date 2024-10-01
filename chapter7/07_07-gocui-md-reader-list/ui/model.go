package ui

import (
	"github.com/awesome-gocui/gocui"
)

type mdfile struct {
	name       string
	content    string
	isSelected bool
}

var (
	mdfiles          = []mdfile{}
	scrollPercentage int
	selectedIndex    int
)

func NewMDReader() error {
	g, err := gocui.NewGui(gocui.OutputNormal, false)
	if err != nil {
		return err
	}
	defer g.Close()

	g.SetManagerFunc(layout)
	g.Highlight = true
	g.SelFrameColor = gocui.ColorGreen

	// load mdfiles slice
	maxX, _ := g.Size()
	mdfiles, err = getMarkdownFiles("docs", maxX)
	if err != nil {
		return err
	}

	mdfiles[selectedIndex].isSelected = true

	err = setKeyBindings(g)
	if err != nil {
		return err
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}

	return nil
}
