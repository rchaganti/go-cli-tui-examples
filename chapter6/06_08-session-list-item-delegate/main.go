package main

import (
	"encoding/json"
	"log"
	"os"
	"session-list/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	sessions := []ui.Session{}

	//read data/history.json and marshal it to a slice of session structs
	f, err := os.ReadFile("data/history.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(f, &sessions)
	if err != nil {
		log.Fatal(err)
	}

	//initialize a new sessionModel with the session list
	model := ui.InitialModel(sessions)

	//initialize a new program with the sessionModel
	p := tea.NewProgram(
		model,
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
