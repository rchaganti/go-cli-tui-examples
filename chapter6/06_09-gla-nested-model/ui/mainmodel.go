package ui

import (
	helper "gla/internal"

	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type state int

const (
	chatView state = iota
	listView
)

type mainModel struct {
	chat     chatModel
	sessions sessionModel
	state    state
	help     help.Model
	keys     keyMap
}

func (m *mainModel) Init() tea.Cmd {
	return nil
}

func (m *mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		sCmd           tea.Cmd
		cCmd           tea.Cmd
		session        helper.Session
		updateChatView bool
	)

	cmds := []tea.Cmd{}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "ctrl+h":
			m.help.ShowAll = !m.help.ShowAll

		case "ctrl+l":
			if m.state == chatView {
				m.state = listView
				m.keys = listKeys
				m.sessions, sCmd = m.sessions.Update(msg)
				cmds = append(cmds, sCmd)
			} else {
				m.state = chatView
				m.keys = chatKeys
				m.chat, cCmd = m.chat.Update(msg)
				cmds = append(cmds, cCmd)
			}

		default:
			if m.state == chatView {
				m.chat, cCmd = m.chat.Update(msg)
				cmds = append(cmds, cCmd)
			} else {
				m.sessions, sCmd = m.sessions.Update(msg)
				cmds = append(cmds, sCmd)
			}
		}

	case sessionSelectedMsg:
		session = msg.session
		updateChatView = true

	case sessionDeletedMsg:
		session = msg.session
		updateChatView = true

	case viewChangedMsg:
		view := msg.view
		m.keys.view = viewState(view)

	default:
		if m.state == chatView {
			m.chat, cCmd = m.chat.Update(msg)
			cmds = append(cmds, cCmd)
		} else {
			m.sessions, sCmd = m.sessions.Update(msg)
			cmds = append(cmds, sCmd)
		}
	}

	if updateChatView {
		m.chat, cCmd = m.chat.Update(msg)
		m.chat.messageHistory = session.Messages
		m.chat.sessionId = session.Id
		m.chat.tokenCount = session.TokenCount
		m.chat.ai.model = session.Model
		m.chat.viewport.SetContent(renderContent(session.Messages))
		m.chat.viewport.GotoBottom()
		cmds = append(cmds, cCmd)
		m.state = chatView
	}

	return m, tea.Batch(cmds...)
}

func (m *mainModel) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Right, lipgloss.JoinHorizontal(
			lipgloss.Right, m.listView(), m.chat.View(),
		), m.helpView(),
	)
}

func InitialModel(apikey string, session helper.Session, sessionStore helper.SessionStore, sessionEntries []helper.Session) *mainModel {
	help := help.New()
	help.Styles.FullKey = helpKeysStyle
	help.Styles.ShortKey = helpKeysStyle
	help.Styles.ShortDesc = helpDescStyle
	help.Styles.FullDesc = helpDescStyle

	return &mainModel{
		chat:     InitialChatModel(apikey, session, sessionStore),
		sessions: InitialSessionModel(sessionEntries, sessionStore),
		state:    chatView,
		help:     help,
		keys:     chatKeys,
	}
}
