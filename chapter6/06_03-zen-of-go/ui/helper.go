package ui

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m model) headerView() string {
	title := titleStyle.Render("The Zen of Go")
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m model) footerView() string {
	info := infoStyle.MarginTop(1).Render(fmt.Sprintf("%d of %d", m.index+1, len(m.content)))
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func (m model) helpView() string {
	return helpStyle("\n  ←/→: Navigate • h: Home •  e: End • q: Quit\n")
}

func getMDContent(path string) (content []string, err error) {
	fileInfo, err := os.ReadDir(path)
	if err != nil {
		log.Fatal("Error reading directory:", err)
	}

	for _, file := range fileInfo {
		fName := fmt.Sprintf("%s%s", path, file.Name())
		if !file.IsDir() && strings.HasSuffix(fName, ".md") {
			text, err := os.ReadFile(fName)
			if err != nil {
				log.Fatal("could not load file:", err)
			}
			content = append(content, string(text))
		}
	}

	return
}
