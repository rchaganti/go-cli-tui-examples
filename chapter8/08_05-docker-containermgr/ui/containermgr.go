package ui

import (
	helper "docker-containermgr/internal"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/docker/cli/cli/command"
	"github.com/muesli/reflow/wordwrap"
)

type state int

const (
	listView state = iota
	tabView
)

type containermgr struct {
	containerList containerListModel
	viewer        viewerModel
	state         state
	help          help.Model
	keys          keyMap
}

func (c containermgr) Init() tea.Cmd {
	return nil
}

func (c containermgr) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		cmds = append(cmds, c.containerList.selectContainer())
	case errMsg:
		cmds = append(cmds, tea.Quit)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			cmds = append(cmds, tea.Quit)
		case "?":
			c.help.ShowAll = !c.help.ShowAll
		case "tab":
			if c.state == tabView {
				c.state = listView
				c.keys = listKeys
				c.containerList.isFocused = true
				c.viewer.isFocused = false
				c.containerList, cmd = c.containerList.Update(msg)
				cmds = append(cmds, cmd)
			} else {
				c.state = tabView
				c.keys = viewerKeys
				c.containerList.isFocused = false
				c.viewer.isFocused = true
				c.viewer, cmd = c.viewer.Update(msg)
				cmds = append(cmds, cmd)
			}
		default:
			if c.state == listView {
				c.containerList, cmd = c.containerList.Update(msg)
				cmds = append(cmds, cmd)
			} else {
				c.viewer, cmd = c.viewer.Update(msg)
				cmds = append(cmds, cmd)
			}
		}
	case containerSelectedMsg:
		var content string
		c.viewer.content.inspect = msg.inspect
		c.viewer.content.logs = msg.logs
		if c.viewer.activeTabIndex == 0 {
			content = c.viewer.content.logs
		} else {
			content = c.viewer.content.inspect
		}

		wrappedContent := wordwrap.String(content, docWidth-clWidth)
		c.viewer.viewer.SetContent(wrappedContent)

		c.viewer, cmd = c.viewer.Update(msg)
		cmds = append(cmds, cmd)
	case containerStoppedMsg:
		var uCmd tea.Cmd
		for i, item := range c.containerList.containers.Items() {
			if p, ok := item.(helper.Container); ok {
				if p.ID == msg.stoppedContainerId {
					p.State = "exited"
					uCmd = c.containerList.containers.SetItem(i, p)
					cmds = append(cmds, uCmd)
					break
				}
			}
		}

		c.containerList, cmd = c.containerList.Update(msg)
		cmds = append(cmds, cmd)
	default:
		if c.state == listView {
			c.containerList, cmd = c.containerList.Update(msg)
			cmds = append(cmds, cmd)
		} else {
			c.viewer, cmd = c.viewer.Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	return c, tea.Batch(cmds...)
}

func (c containermgr) View() string {
	return lipgloss.JoinVertical(lipgloss.Right,
		lipgloss.JoinHorizontal(lipgloss.Top, c.containerList.View(), c.viewer.View()),
		c.helpView(),
	)
}

func NewContainerMgr(c []helper.Container, dockerCli command.Cli) containermgr {
	help := help.New()
	help.Styles.FullKey = helpKeysStyle
	help.Styles.ShortKey = helpKeysStyle
	help.Styles.ShortDesc = helpDescStyle
	help.Styles.FullDesc = helpDescStyle

	return containermgr{
		containerList: NewContainerListModel(c, dockerCli.CurrentContext(), dockerCli),
		viewer:        NewViewerModel(),
		state:         listView,
		help:          help,
		keys:          listKeys,
	}
}
