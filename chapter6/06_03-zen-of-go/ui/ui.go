package ui

import (
	"log"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	content  []string
	index    int
	viewport viewport.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		case "right":
			if m.index < len(m.content)-1 {
				m.index++
			}
		case "left":
			if m.index > 0 {
				m.index--
			}
		case "h":
			m.index = 0
		case "e":
			m.index = len(m.content) - 1
		}
	}

	return m, nil
}

func (m model) View() string {
	style := glamour.DarkStyleConfig

	renderer, err := glamour.NewTermRenderer(
		glamour.WithStyles(style),
		glamour.WithPreservedNewLines(),
		glamour.WithWordWrap(65),
	)

	if err != nil {
		log.Fatal(err)
	}

	str, err := renderer.Render(m.content[m.index])
	if err != nil {
		log.Fatal(err)
	}

	m.viewport.SetContent(str)
	return m.headerView() + m.viewport.View() + m.footerView() + m.helpView()
}

func IntialModel(path string) model {
	content, err := getMDContent(path)
	if err != nil {
		log.Fatal(err)
	}

	vp := viewport.New(78, 22)
	vp.Style = lipgloss.NewStyle().Padding(5)

	return model{
		content:  content,
		index:    0,
		viewport: vp,
	}
}
