package ui

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Navigate key.Binding
	Enter    key.Binding
	Refresh  key.Binding
	Quit     key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Navigate, k.Enter, k.Refresh, k.Quit,
	}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Navigate, k.Enter},
		{k.Refresh, k.Quit},
	}
}

var listKeys = keyMap{
	Navigate: key.NewBinding(
		key.WithKeys("down", "up"),
		key.WithHelp("↓/↑", "down/up"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("⏎", "switch context"),
	),
	Refresh: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "refresh context"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
}
