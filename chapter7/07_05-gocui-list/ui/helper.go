package ui

import (
	"fmt"

	"github.com/awesome-gocui/gocui"
)

func moveViewCursor(v *gocui.View, ty int) error {
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

func toggleNamespaceSelection(l string) error {
	for i, namespace := range namespaces {
		if namespace.name == l {
			if !namespace.isSelected {
				namespaces[i].isSelected = true
			}
		} else {
			namespaces[i].isSelected = false
		}
	}

	return nil
}

func viewNamespacesRefreshList(g *gocui.Gui) {
	var str string
	var y int

	g.Update(func(g *gocui.Gui) error {
		v, err := g.View("namespaces")
		if err != nil {
			return err
		}

		v.Clear()

		if len(namespaces) > 0 {
			for i, namespace := range namespaces {
				if namespace.isSelected {
					str = fmt.Sprintf("*%-20s", namespace.name)
					y = i
				} else {
					str = fmt.Sprintf(" %-20s", namespace.name)
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
