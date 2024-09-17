package ui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Session struct {
	SessionId  string `json:"sessionid"`
	Timestamp  string `json:"timestamp"`
	TokenCount int    `'json:"tokencount"`
	Model      string `json:"model"`
	IsSelected bool   `json:"isselected"`
}

func (s Session) FilterValue() string {
	return s.SessionId
}

func (s Session) Title() string {
	return s.SessionId
}

func (s Session) Description() string {
	return s.Timestamp
}

type sessionModel struct {
	sessionList list.Model
	help        help.Model
	keys        KeyMap
}

func (m sessionModel) Init() tea.Cmd {
	return nil
}

func (m sessionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		uCmd tea.Cmd
		sCmd tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.sessionList.SetSize(width, height)

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "ctrl+h":
			m.help.ShowAll = !m.help.ShowAll

		case "delete":
			index := m.sessionList.Index()

			// delete the selected item
			s, ok := m.sessionList.SelectedItem().(Session)
			m.sessionList.RemoveItem(index)

			if ok && s.IsSelected {
				if len(m.sessionList.Items()) > 0 {
					item := m.sessionList.Items()[0].(Session)
					item.IsSelected = true
					sCmd = m.sessionList.SetItem(0, item)
					return m, sCmd
				}
			}

		case "enter":
			currentIndex := m.sessionList.Index()
			if s, ok := m.sessionList.SelectedItem().(Session); ok {
				for i, item := range m.sessionList.Items() {
					if p, ok := item.(Session); ok {
						if p.SessionId != s.SessionId {
							p.IsSelected = false
							m.sessionList.SetItem(i, p)
						}
					}
				}

				// set the new selected item
				s.IsSelected = true
				sCmd = m.sessionList.SetItem(currentIndex, s)

				return m, sCmd
			}

		default:
			m.sessionList, uCmd = m.sessionList.Update(msg)
			return m, uCmd
		}
	}

	m.sessionList, uCmd = m.sessionList.Update(msg)
	return m, uCmd
}

func (m sessionModel) View() string {
	helpView := m.help.View(m.keys)

	return lipgloss.JoinVertical(lipgloss.Left, listStyle.Render(m.sessionList.View()), helpView)
}

func InitialModel(sessions []Session) sessionModel {
	listItems := []list.Item{}
	for _, s := range sessions {
		listItems = append(listItems, s)
	}

	list := list.New(listItems, itemDelegate{}, 0, 0)

	list.SetShowStatusBar(false)
	list.SetShowHelp(false)

	list.Title = "Conversations"

	help := help.New()
	help.Styles.FullKey = helpKeysStyle
	help.Styles.ShortKey = helpKeysStyle
	help.Styles.ShortDesc = helpDescStyle
	help.Styles.FullDesc = helpDescStyle

	return sessionModel{
		sessionList: list,
		help:        help,
		keys:        sessionKeys,
	}
}
