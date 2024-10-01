package ui

import (
	"strings"

	"github.com/awesome-gocui/gocui"
)

func setKeybindings(g *gocui.Gui) {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		panic(err)
	}

	if err := g.SetKeybinding("namespaces", gocui.KeyArrowDown, gocui.ModNone, moveViewCursorDown); err != nil {
		panic(err)
	}

	if err := g.SetKeybinding("namespaces", gocui.KeyArrowUp, gocui.ModNone, moveViewCursorUp); err != nil {
		panic(err)
	}

	if err := g.SetKeybinding("namespaces", gocui.KeyEnter, gocui.ModNone, selectNamespace); err != nil {
		panic(err)
	}

	if err := g.SetKeybinding("namespaces", gocui.KeyCtrlR, gocui.ModNone, renameView); err != nil {
		panic(err)
	}

	if err := g.SetKeybinding("rename", gocui.KeyEnter, gocui.ModNone, finalizeRename); err != nil {
		panic(err)
	}

	if err := g.SetKeybinding("rename", gocui.KeyEsc, gocui.ModNone, deleteRename); err != nil {
		panic(err)
	}
}

func selectNamespace(g *gocui.Gui, v *gocui.View) error {
	_, cy := v.Cursor()
	l, err := v.Line(cy)
	if err != nil {
		return nil
	}

	l = strings.Trim(l, " ")
	l = strings.TrimPrefix(l, "*")
	if err := toggleNamespaceSelection(l); err != nil {
		return err
	}

	go viewNamespacesRefreshList(g)
	return nil
}

func finalizeRename(g *gocui.Gui, v *gocui.View) error {
	var nvl, l string
	var err error

	// get the string user wants to rename
	nv, err := g.View("namespaces")
	if err != nil {
		return err
	}

	_, nvcy := nv.Cursor()
	if nvl, err = nv.Line(nvcy); err != nil {
		nvl = ""
		return nil
	}

	// remove additional spaces and *
	nvl = strings.Trim(nvl, " ")
	nvl = strings.TrimPrefix(nvl, "*")

	_, cy := v.Cursor()
	if l, err = v.Line(cy); err != nil {
		l = ""
		return nil
	}

	// rename the string in namespaces slice
	for i, namespace := range namespaces {
		if namespace.name == nvl {
			namespaces[i].name = l
		}
	}

	// delete the rename view
	deleteRename(g, v)
	return nil
}

func deleteRename(g *gocui.Gui, v *gocui.View) error {
	err := g.DeleteView("rename")
	go viewNamespacesRefreshList(g)
	g.SetCurrentView("namespaces")
	g.Cursor = false
	if err != nil {
		return err
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func moveViewCursorDown(_ *gocui.Gui, v *gocui.View) error {
	err := moveViewCursor(v, 1)
	if err != nil {
		return err
	}
	return nil
}

func moveViewCursorUp(_ *gocui.Gui, v *gocui.View) error {
	err := moveViewCursor(v, -1)
	if err != nil {
		return err
	}
	return nil
}
