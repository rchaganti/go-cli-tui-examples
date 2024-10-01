package ui

import "github.com/awesome-gocui/gocui"

var (
	scrollPercentage int
	mdContent        string
)

func NewMDReader() error {
	g, err := gocui.NewGui(gocui.OutputNormal, false)
	if err != nil {
		return err
	}
	defer g.Close()

	g.SetManagerFunc(layout)
	g.SelFrameColor = gocui.ColorGreen

	err = setKeyBindings(g)
	if err != nil {
		return err
	}

	maxX, _ := g.Size()
	filename := "docs/computer-history.md"
	mdContent = convertMdtoTerm(filename, maxX)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		return err
	}

	return nil
}
