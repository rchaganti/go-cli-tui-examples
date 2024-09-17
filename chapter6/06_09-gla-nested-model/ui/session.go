package ui

import (
	helper "gla/internal"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type sessionModel struct {
	sessionList  list.Model
	sessionStore helper.SessionStore
}

func (m sessionModel) Init() tea.Cmd {
	return nil
}

func (m sessionModel) Update(msg tea.Msg) (sessionModel, tea.Cmd) {
	var (
		uCmd tea.Cmd
		sCmd tea.Cmd
	)

	cmds := []tea.Cmd{}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.sessionList.SetSize(slWidth, slHeight)

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			cmds = append(cmds, tea.Quit)

		case "delete":
			index := m.sessionList.Index()

			// delete the selected item
			s, ok := m.sessionList.SelectedItem().(helper.Session)
			m.sessionList.RemoveItem(index)

			if ok && s.IsSelected {
				if len(m.sessionList.Items()) > 0 {
					item := m.sessionList.Items()[0].(helper.Session)
					item.IsSelected = true
					sCmd = m.sessionList.SetItem(0, item)
					cmds = append(cmds, sCmd)
				}
			}
			cmds = append(cmds, m.deleteSession(s))

		case "enter":
			currentIndex := m.sessionList.Index()
			if s, ok := m.sessionList.SelectedItem().(helper.Session); ok {
				for i, item := range m.sessionList.Items() {
					if p, ok := item.(helper.Session); ok {
						if p.Id != s.Id {
							p.IsSelected = false
							m.sessionList.SetItem(i, p)
						}
					}
				}

				// set the new selected item
				s.IsSelected = true
				sCmd = m.sessionList.SetItem(currentIndex, s)
				cmds = append(cmds, sCmd, m.selectSession(s))
			}

		default:
			m.sessionList, uCmd = m.sessionList.Update(msg)
			return m, uCmd
		}
	}

	m.sessionList, uCmd = m.sessionList.Update(msg)
	cmds = append(cmds, uCmd)

	return m, tea.Batch(cmds...)
}

func (m sessionModel) View() string {
	return m.sessionList.View()
}

func InitialSessionModel(sessions []helper.Session, sessionStore helper.SessionStore) sessionModel {
	var (
		selectedIndex int
		listItems     []list.Item
	)

	listItems = []list.Item{}
	for i, s := range sessions {
		if s.IsSelected {
			selectedIndex = i
		}
		listItems = append(listItems, s)
	}

	list := list.New(listItems, itemDelegate{}, slWidth, slHeight-8)

	list.SetShowStatusBar(false)
	list.SetShowHelp(false)

	list.Title = "Conversations"
	list.Select(selectedIndex)

	return sessionModel{
		sessionList:  list,
		sessionStore: sessionStore,
	}
}
