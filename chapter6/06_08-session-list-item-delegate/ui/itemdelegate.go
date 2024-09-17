package ui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type itemDelegate struct{}

func (d itemDelegate) Height() int {
	return 1
}

func (d itemDelegate) Spacing() int {
	return 0
}

func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd {
	return nil
}

func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	var (
		title, desc string
	)
	s, ok := listItem.(Session)
	if ok {
		title = s.Title()
		desc = s.Description()
	} else {
		return
	}

	selectedIndex := m.Index()

	if s.IsSelected {
		fmt.Fprintf(w, "%s\n%s\n", selectedTitle.Render(title), selectedDesc.Render(desc))
	} else if index == selectedIndex {
		fmt.Fprintf(w, "%s\n%s\n", normalTitle.Render(title), normalDesc.Render(desc))
	} else {
		fmt.Fprintf(w, "%s\n%s\n", dimmedTitle.Render(title), dimmedDesc.Render(desc))
	}
}
