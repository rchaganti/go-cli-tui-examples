package ui

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/wordwrap"
)

type tabContent struct {
	logs    string
	inspect string
}

type viewerModel struct {
	width          int
	tabTitles      []string
	activeTabIndex int
	viewer         viewport.Model
	content        tabContent
	isFocused      bool
}

func (v viewerModel) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (v viewerModel) Update(msg tea.Msg) (viewerModel, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	var updateViewer bool

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		v.width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "right":
			v.activeTabIndex = (v.activeTabIndex + 1) % len(v.tabTitles)
			updateViewer = true
		case "left":
			v.activeTabIndex = (v.activeTabIndex - 1 + len(v.tabTitles)) % len(v.tabTitles)
			updateViewer = true
		case "up", "pgup":
			v.viewer, cmd = v.viewer.Update(msg)
			cmds = append(cmds, cmd)
		case "down", "pgdown":
			v.viewer, cmd = v.viewer.Update(msg)
			cmds = append(cmds, cmd)
		case "ctrl+c", "q":
			cmds = append(cmds, tea.Quit)
		}
	}

	if updateViewer {
		var content string
		if v.activeTabIndex == 0 {
			content = v.content.logs
		} else {
			content = v.content.inspect
		}

		wrappedContent := wordwrap.String(content, docWidth-clWidth-2)
		v.viewer.SetContent(wrappedContent)

		v.viewer, cmd = v.viewer.Update(msg)
		cmds = append(cmds, cmd)
	}

	return v, tea.Batch(cmds...)
}

func (m viewerModel) View() string {
	return m.viewerView()
}

func NewViewerModel() viewerModel {
	vp := viewport.Model{
		Width:  docWidth - clWidth - 2,
		Height: docHeight - 2,
	}

	return viewerModel{
		width:     docWidth - clWidth,
		tabTitles: []string{"Logs", "Inspect"},
		viewer:    vp,
		content:   tabContent{},
	}
}
