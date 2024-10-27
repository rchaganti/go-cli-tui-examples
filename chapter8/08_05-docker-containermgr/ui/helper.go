package ui

import (
	helper "docker-containermgr/internal"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type containerSelectedMsg struct {
	logs    string
	inspect string
}

type containerStoppedMsg struct {
	stoppedContainerId string
}

type errMsg struct{ Err error }

// view helpers
func (l containerListModel) listView() string {
	if !l.isFocused {
		return listFaintStyle.Render(l.containers.View())
	}
	return listFocusedStyle.Render(l.containers.View())
}

func (v viewerModel) tabsView() string {
	var tabRenders []string
	var borderColor string
	if v.isFocused {
		borderColor = focusedBorder
	} else {
		borderColor = faintBorder
	}

	for i, title := range v.tabTitles {
		if i == v.activeTabIndex {
			tabRenders = append(tabRenders, activeTabStyle.BorderForeground(lipgloss.Color(borderColor)).Render(title))
		} else {
			tabRenders = append(tabRenders, inactiveTabStyle.BorderForeground(lipgloss.Color(borderColor)).Render(title))
		}
	}

	tabs := lipgloss.JoinHorizontal(lipgloss.Top, tabRenders...)

	tabsWidth := lipgloss.Width(tabs)
	remainingWidth := v.width - tabsWidth

	if remainingWidth < 0 {
		remainingWidth = 0
	}

	extendLine := lipgloss.NewStyle().
		BorderStyle(extendBorder).
		//BorderBottom(true).
		BorderForeground(lipgloss.Color(borderColor)).
		Width(remainingWidth - 1).Render()

	row := lipgloss.JoinHorizontal(lipgloss.Bottom, tabs, extendLine)

	return row
}

func (v viewerModel) viewerView() string {
	tabs := v.tabsView()
	if v.isFocused {
		return lipgloss.JoinVertical(lipgloss.Left, tabs, viewerStyle.BorderForeground(lipgloss.Color(focusedBorder)).Render(v.viewer.View()))
	} else {
		return lipgloss.JoinVertical(lipgloss.Left, tabs, viewerStyle.BorderForeground(lipgloss.Color(faintBorder)).Render(v.viewer.View()))
	}
}

func (c containermgr) helpView() string {
	view := c.help.View(c.keys)
	return helpStyle.Render(view)
}

// bubbletea commands
func (l containerListModel) selectContainer() tea.Cmd {
	return func() tea.Msg {
		selectedContainer, ok := l.containers.SelectedItem().(helper.Container)
		if ok {
			cid := selectedContainer.ID

			// get logs
			logs, err := helper.GetContainerLogs(l.dockerCli, cid)
			if err != nil {
				return errMsg{Err: err}
			}

			// get inspect
			inspect, err := helper.GetContainerInspect(l.dockerCli, cid)
			if err != nil {
				return errMsg{Err: err}
			}

			return containerSelectedMsg{logs: logs, inspect: inspect}
		} else {
			return errMsg{Err: fmt.Errorf("failed to select container")}
		}

	}
}

func (l containerListModel) stopContainer() tea.Cmd {
	return func() tea.Msg {
		selectedContainer, ok := l.containers.SelectedItem().(helper.Container)
		if ok {
			err := helper.StopContainer(l.dockerCli, selectedContainer.ID)
			if err != nil {
				return errMsg{Err: err}
			}

			return containerStoppedMsg{
				stoppedContainerId: selectedContainer.ID,
			}
		}
		return errMsg{Err: fmt.Errorf("failed to stop container")}
	}
}
