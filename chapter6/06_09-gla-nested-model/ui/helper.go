package ui

import (
	"fmt"
	helper "gla/internal"
	"gla/pkg/gai"
	"strings"

	markdown "github.com/MichaelMure/go-term-markdown"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/generative-ai-go/genai"
)

type sessionSavedMsg struct {
	session helper.Session
}

type sessionDeletedMsg struct {
	session helper.Session
}

type sessionSelectedMsg struct {
	session helper.Session
}

func (m mainModel) helpView() string {
	view := m.help.View(m.keys)
	return helpStyle.Render(view)
}

func (m chatModel) headerView() string {
	var title, headerText, line, leftEdge, view string
	headerText = fmt.Sprintf("Gemini Learning Assistant ♦ %s", m.ai.model)

	if m.viewState == chatinput {
		title = titleBlurStyle.Render(headerText)
		line = strings.Repeat(blurStyle.Render(l), max(0, m.viewport.Width-lipgloss.Width(title)))
		leftEdge = blurStyle.Render(le)
	} else {
		title = titleStyle.Render(headerText)
		line = strings.Repeat(focusStyle.Render(l), max(0, m.viewport.Width-lipgloss.Width(title)))
		leftEdge = focusStyle.Render(le)
	}

	view = lipgloss.JoinHorizontal(lipgloss.Center, title, line, leftEdge)

	return view
}

func (m chatModel) footerView() string {
	var info, line, rightEdge, view string

	footerString := fmt.Sprintf("Tokens: %d", m.tokenCount)
	if m.viewState == chatinput {
		info = infoBlurStyle.Render(footerString)
		line = strings.Repeat(blurStyle.Render(l), max(0, m.viewport.Width-lipgloss.Width(info)))
		rightEdge = blurStyle.Render(re)
	} else {
		info = infoStyle.Render(footerString)
		line = strings.Repeat(focusStyle.Render(l), max(0, m.viewport.Width-lipgloss.Width(info)))
		rightEdge = focusStyle.Render(re)
	}

	view = lipgloss.JoinHorizontal(lipgloss.Center, rightEdge, line, info)

	return view
}

func (m chatModel) textAreaView() string {
	if m.viewState == chatinput {
		return baseStyle.Render(m.textarea.View())
	} else {
		return baseBlurStyle.Render(m.textarea.View())
	}
}

func (m chatModel) spinnerView() string {
	spContent := fmt.Sprintf("Churning the AI universe %s", m.spinner.View())
	if m.viewState == chatinput {
		return baseStyle.Render(spContent)
	} else {
		return baseBlurStyle.Render(spContent)
	}
}

func (m mainModel) listView() string {
	if m.state == chatView {
		return listFaintStyle.Render(m.sessions.View())
	} else {
		return listFocusedStyle.Render(m.sessions.View())
	}
}

func (m chatModel) generateResponse(prompt string) tea.Cmd {
	return func() tea.Msg {
		var res *genai.GenerateContentResponse
		res, err := gai.GetResponse(prompt, m.messageHistory, m.ai.apiKey, m.ai.model)
		if err != nil {
			return errMsg{Err: err}
		}

		result := gai.PrintResponse(res)
		return responseMsg{response: result, tokenCount: int(res.UsageMetadata.TotalTokenCount)}
	}
}

func (m chatModel) addMessageToSession(message gai.MessageContent) tea.Cmd {
	return func() tea.Msg {
		session, err := m.sessionStore.AddMessage(m.sessionId, message, m.tokenCount)
		if err != nil {
			return errMsg{Err: err}
		}
		return sessionSavedMsg{session: session}
	}
}

func (m sessionModel) deleteSession(session helper.Session) tea.Cmd {
	return func() tea.Msg {
		session, err := m.sessionStore.DeleteSession(session.Id)
		if err != nil {
			return errMsg{Err: err}
		}
		return sessionDeletedMsg{session: session}
	}
}

func (m sessionModel) selectSession(session helper.Session) tea.Cmd {
	return func() tea.Msg {
		session, err := m.sessionStore.SelectSession(session.Id)
		if err != nil {
			return errMsg{Err: err}
		}
		return sessionSelectedMsg{session: session}
	}
}

func (m chatModel) changeView() tea.Cmd {
	return func() tea.Msg {
		if m.viewState == chatinput {
			return viewChangedMsg{view: chatinput}
		}
		return viewChangedMsg{view: chatview}
	}
}

func renderContent(messageHistory []gai.MessageContent) string {
	var messages []string
	for _, mesg := range messageHistory {
		if mesg.Role == "user" {
			messages = append(messages, userMessageStyle.Render("👤: "+mesg.Message))
		} else {
			result := string(markdown.Render(mesg.Message, vpWidth, 1))
			messages = append(messages, geminiMessageStyle.Render("🤖: "+strings.Trim(result, " ")))
		}
	}

	return strings.Join(messages, "\n")
}
