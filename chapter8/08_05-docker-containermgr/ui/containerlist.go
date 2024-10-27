package ui

import (
	helper "docker-containermgr/internal"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/docker/cli/cli/command"
)

type containerListModel struct {
	containers     list.Model
	currentContext string
	dockerCli      command.Cli
	isFocused      bool
}

func (l containerListModel) Init() tea.Cmd {
	return nil
}

func (l containerListModel) Update(msg tea.Msg) (containerListModel, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			cmds = append(cmds, tea.Quit)
		case "s":
			if s, ok := l.containers.SelectedItem().(helper.Container); ok {
				if s.State == "running" {
					cmds = append(cmds, l.stopContainer())
				}
			}
		case "enter":
			currentIndex := l.containers.Index()
			if s, ok := l.containers.SelectedItem().(helper.Container); ok {
				for i, item := range l.containers.Items() {
					if p, ok := item.(helper.Container); ok {
						if p.ID != s.ID {
							p.IsSelected = false
							l.containers.SetItem(i, p)
						}
					}
				}

				// set the new selected item
				s.IsSelected = true
				cmd = l.containers.SetItem(currentIndex, s)
				cmds = append(cmds, cmd, l.selectContainer())
			}
		default:
			l.containers, cmd = l.containers.Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	return l, tea.Batch(cmds...)
}

func (l containerListModel) View() string {
	return l.listView()
}

func NewContainerListModel(containers []helper.Container, currentContext string, dockerCli command.Cli) containerListModel {
	var (
		listItems []list.Item
	)

	listItems = []list.Item{}
	for _, c := range containers {
		listItems = append(listItems, c)
	}

	list := list.New(listItems, itemDelegate{}, clWidth, clHeight-8)

	list.SetShowStatusBar(false)
	list.SetShowHelp(false)
	list.SetFilteringEnabled(false)

	list.Title = currentContext
	list.Select(0)

	return containerListModel{
		containers:     list,
		currentContext: currentContext,
		dockerCli:      dockerCli,
		isFocused:      true,
	}
}
