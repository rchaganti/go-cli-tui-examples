package ui

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/awesome-gocui/gocui"
)

func convertMdtoTerm(filename string, maxX int) string {
	source, err := os.ReadFile(filename)
	if err != nil {
		log.Print(err)
	}

	result := markdown.Render(string(source), maxX-8, 6)
	return string(result)
}

func getMarkdownFiles(dir string, maxX int) ([]mdfile, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".md" {
			mdContent := convertMdtoTerm(filepath.Join(dir, file.Name()), maxX)
			mdfiles = append(mdfiles, mdfile{name: file.Name(), content: mdContent})
		}
	}

	return mdfiles, nil
}

func updateScrollPercentage(g *gocui.Gui, v *gocui.View) {
	_, oy := v.Origin()
	_, height := v.Size()

	totalLines := len(v.ViewBufferLines())

	scrollPercentage = int((float64(oy) / float64(totalLines-height)) * 100)

	v.Title = fmt.Sprintf(" Markdown Reader [ %d%% ]", scrollPercentage)
	g.Update(func(g *gocui.Gui) error { return nil })
}

func moveListCursor(v *gocui.View, ty int) error {
	l, err := getViewLine(v, ty)
	if err != nil {
		return err
	}

	if l != "" {
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy+ty); err != nil {
			return err
		}
	} else {
		return nil
	}
	return nil
}

func getViewLine(v *gocui.View, ty int) (string, error) {
	var l string
	var err error

	_, cy := v.Cursor()

	if ny := cy + ty; ny >= 0 {
		if l, err = v.Line(ny); err != nil {
			l = ""
		}
	}

	return l, err
}

func toggleFileSelection(l string) error {
	for i, mdfile := range mdfiles {
		if mdfile.name == l {
			if !mdfile.isSelected {
				mdfiles[i].isSelected = true
				selectedIndex = i
			}
		} else {
			mdfiles[i].isSelected = false
		}
	}

	return nil
}

func refreshList(g *gocui.Gui) {
	var str string
	var y int

	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("list")
		if err != nil {
			return err
		}

		v.Clear()

		if len(mdfiles) > 0 {
			for i, mdfile := range mdfiles {
				if mdfile.isSelected {
					str = fmt.Sprintf("*%s", mdfile.name)
					y = i
				} else {
					str = fmt.Sprintf(" %s", mdfile.name)
				}
				fmt.Fprintln(v, str)
			}
		}

		x, _ := v.Cursor()
		if err := v.SetCursor(x, y); err != nil {
			return err
		}

		return nil
	})
}

func refreshContent(g *gocui.Gui) {
	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("reader")
		if err != nil {
			return err
		}

		v.Clear()

		fmt.Fprintln(v, mdfiles[selectedIndex].content)

		return nil
	})
}

func refreshStatusBar(g *gocui.Gui) {
	g.Update(func(g *gocui.Gui) error {
		var str string

		v, err := g.View("status")
		if err != nil {
			return err
		}

		v.Clear()
		cv := g.CurrentView()
		if cv.Name() == "list" {
			str = "Enter: Select File | ↑ ↓: Select File | Tab: Switch View | Ctrl-C: Exit"
		} else {
			str = "↑ ↓: Scroll Content | Tab: Switch View | Ctrl-C: Exit"
		}

		fmt.Fprint(v, str)
		return nil
	})
}
