package ui

import (
	"strings"

	helper "gla/internal"
	"gla/pkg/gai"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type viewState int

const (
	chatview viewState = iota
	chatinput
	listview
)

type responseMsg struct {
	response   string
	tokenCount int
}

type viewChangedMsg struct {
	view viewState
}

type errMsg struct{ Err error }

type gemini struct {
	apiKey string
	model  string
}

type chatModel struct {
	ai             gemini
	viewport       viewport.Model
	textarea       textarea.Model
	spinner        spinner.Model
	messageHistory []gai.MessageContent
	tokenCount     int
	viewState      viewState
	responding     bool
	err            error
	sessionStore   helper.SessionStore
	sessionId      string
}

func (m chatModel) Init() tea.Cmd {
	return nil
}

func (m chatModel) Update(msg tea.Msg) (chatModel, tea.Cmd) {
	var (
		vpCmd tea.Cmd
		taCmd tea.Cmd
		spCmd tea.Cmd
	)

	var cmds = []tea.Cmd{}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.viewport.GotoBottom()

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "enter":
			prompt := m.textarea.Value()
			if (strings.Trim(prompt, " ")) == "" {
				return m, nil
			}

			message := gai.MessageContent{
				Message: prompt,
				Role:    "user",
			}
			m.messageHistory = append(m.messageHistory, message)

			m.responding = true

			m.viewport, vpCmd = m.viewport.Update(msg)
			m.viewport.SetContent(renderContent(m.messageHistory))
			m.viewport.GotoBottom()

			m.spinner, spCmd = m.spinner.Update(msg)

			cmds = append(cmds, vpCmd, spCmd, m.spinner.Tick, m.addMessageToSession(message), m.generateResponse(prompt))
			m.viewState = chatview

		case "tab":
			if m.viewState == chatview {
				m.viewState = chatinput
				cmds = append(cmds, m.changeView())
			} else {
				m.viewState = chatview
				cmds = append(cmds, m.changeView())
			}
		default:
			if m.viewState == chatinput {
				m.textarea, taCmd = m.textarea.Update(msg)
				cmds = append(cmds, taCmd, textarea.Blink)
			} else {
				m.viewport, vpCmd = m.viewport.Update(msg)
				cmds = append(cmds, vpCmd)
			}
		}

	case responseMsg:
		m.tokenCount = msg.tokenCount
		message := gai.MessageContent{
			Message: msg.response,
			Role:    "model",
		}
		m.messageHistory = append(m.messageHistory, message)

		m.responding = false

		m.viewport.SetContent(renderContent(m.messageHistory))
		m.viewport.GotoBottom()

		m.textarea, taCmd = m.textarea.Update(msg)
		m.textarea.Reset()
		cmds = append(cmds, taCmd, m.addMessageToSession(message), textarea.Blink)
		m.viewState = chatinput

	case errMsg:
		m.err = msg.Err
		return m, tea.Quit

	case cursor.BlinkMsg:
		m.textarea, taCmd = m.textarea.Update(msg)
		cmds = append(cmds, taCmd)

	case spinner.TickMsg:
		m.spinner, spCmd = m.spinner.Update(msg)
		cmds = append(cmds, spCmd)

	}

	return m, tea.Batch(cmds...)
}

func (m chatModel) View() string {
	var view, taView, spView string
	taView = m.textAreaView()
	spView = m.spinnerView()

	if !m.responding {
		view = lipgloss.JoinVertical(lipgloss.Top, m.headerView(), m.viewport.View(), m.footerView(), taView)
	} else {
		view = lipgloss.JoinVertical(lipgloss.Top, m.headerView(), m.viewport.View(), m.footerView(), spView)
	}

	return view
}

func InitialChatModel(apikey string, session helper.Session, sessionStore helper.SessionStore) chatModel {
	// text area model
	ta := textarea.New()
	ta.Placeholder = "Ask me anything..."
	ta.FocusedStyle.CursorLine = cursorStyle
	ta.Prompt = "┃ "
	ta.ShowLineNumbers = false
	ta.KeyMap.InsertNewline.SetEnabled(false)
	ta.SetWidth(taWidth - 1)
	ta.SetHeight(taHeight)
	ta.Focus()

	// View port model
	vp := viewport.New(vpWidth, vpHeight)
	vp.SetContent(renderContent(session.Messages))

	sp := spinner.New()
	sp.Spinner = spinner.Hamburger

	return chatModel{
		ai: gemini{
			apiKey: apikey,
			model:  session.Model,
		},
		viewport:       vp,
		messageHistory: session.Messages,
		tokenCount:     session.TokenCount,
		textarea:       ta,
		spinner:        sp,
		viewState:      chatinput,
		err:            nil,
		responding:     false,
		sessionId:      session.Id,
		sessionStore:   sessionStore,
	}
}
