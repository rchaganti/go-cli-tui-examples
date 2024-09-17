package ui

import (
	"strings"

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
)

type responseMsg struct {
	response   string
	tokenCount int
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
}

func (m chatModel) Init() tea.Cmd {
	return nil
}

func (m chatModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		vpCmd tea.Cmd
		taCmd tea.Cmd
		spCmd tea.Cmd
	)

	var cmds = []tea.Cmd{}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "enter":
			prompt := m.textarea.Value()
			if (strings.Trim(prompt, " ")) == "" {
				return m, nil
			}

			m.messageHistory = append(m.messageHistory, gai.MessageContent{
				Message: prompt,
				Role:    "user",
			})

			m.responding = true

			m.viewport, vpCmd = m.viewport.Update(msg)
			m.viewport.SetContent(renderContent(m.messageHistory))
			m.viewport.GotoBottom()

			m.spinner, spCmd = m.spinner.Update(msg)

			cmds = append(cmds, vpCmd, spCmd, m.spinner.Tick, m.generateResponse(prompt))
			m.viewState = chatview

		case "tab":
			if m.viewState == chatview {
				m.viewState = chatinput
			} else {
				m.viewState = chatview
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
		m.messageHistory = append(m.messageHistory, gai.MessageContent{
			Message: string(msg.response),
			Role:    "model",
		})

		m.responding = false

		m.viewport.SetContent(renderContent(m.messageHistory))
		m.viewport.GotoBottom()

		m.textarea, taCmd = m.textarea.Update(msg)
		m.textarea.Reset()
		cmds = append(cmds, taCmd, textarea.Blink)

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
		view = lipgloss.JoinVertical(1, m.headerView(), m.viewport.View(), m.footerView(), taView)
	} else {
		view = lipgloss.JoinVertical(lipgloss.Top, m.headerView(), m.viewport.View(), m.footerView(), spView)
	}

	return view
}

func InitialChatModel(apikey, model string, messageHistory []gai.MessageContent) chatModel {
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
	vp.SetContent(renderContent(messageHistory))

	sp := spinner.New()
	sp.Spinner = spinner.Hamburger

	return chatModel{
		ai: gemini{
			apiKey: apikey,
			model:  model,
		},
		viewport:       vp,
		messageHistory: messageHistory,
		textarea:       ta,
		spinner:        sp,
		viewState:      chatinput,
		err:            nil,
		responding:     false,
	}
}
