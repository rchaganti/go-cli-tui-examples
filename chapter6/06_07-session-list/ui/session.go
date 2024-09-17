package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Session struct {
	SessionId  string `json:"id"`
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
}

func (m sessionModel) Init() tea.Cmd {
	return nil
}

func (m sessionModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		uCmd tea.Cmd
		sCmd tea.Cmd
	)

	m.sessionList, uCmd = m.sessionList.Update(msg)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.sessionList.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "enter":
			s, ok := m.sessionList.SelectedItem().(Session)
			if ok {
				sCmd = m.sessionList.NewStatusMessage(statusMessageStyle("Selection changed :" + string(s.SessionId)))
				return m, sCmd
			}
		}
	}

	return m, uCmd
}

func (m sessionModel) View() string {
	return listStyle.Render(m.sessionList.View())
}

func InitialModel(sessions []Session) sessionModel {
	listItems := []list.Item{}
	for _, s := range sessions {
		listItems = append(listItems, s)
	}

	list := list.New(listItems, list.NewDefaultDelegate(), 20, 30)
	return sessionModel{
		sessionList: list,
	}
}
