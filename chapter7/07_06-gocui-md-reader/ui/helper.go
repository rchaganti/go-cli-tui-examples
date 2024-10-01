package ui

import (
	"fmt"
	"log"
	"os"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/awesome-gocui/gocui"
)

func convertMdtoTerm(filename string, maxX int) string {
	source, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	result := markdown.Render(string(source), maxX-8, 6)
	return string(result)
}

func updateScrollPercentage(g *gocui.Gui, v *gocui.View) {
	_, oy := v.Origin()
	_, height := v.Size()

	totalLines := len(v.ViewBufferLines())

	scrollPercentage = int((float64(oy) / float64(totalLines-height)) * 100)

	v.Title = fmt.Sprintf(" Markdown Reader [ %d%% ]", scrollPercentage)
	g.Update(func(g *gocui.Gui) error { return nil })
}
