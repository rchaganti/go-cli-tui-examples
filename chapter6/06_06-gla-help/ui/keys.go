package ui

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	view  viewState
	Up    key.Binding
	Down  key.Binding
	Tab   key.Binding
	Enter key.Binding
	Help  key.Binding
	Quit  key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	if k.view == chatinput {
		return []key.Binding{
			k.Tab, k.Enter, k.Help, k.Quit,
		}
	} else {
		return []key.Binding{
			k.Tab, k.Up, k.Down, k.Help, k.Quit,
		}
	}
}

func (k keyMap) FullHelp() [][]key.Binding {
	if k.view == chatinput {
		return [][]key.Binding{
			{k.Enter, k.Up, k.Down},
			{k.Help, k.Quit},
		}
	} else {
		return [][]key.Binding{
			{k.Tab, k.Up, k.Down},
			{k.Help, k.Quit},
		}
	}
}

var chatKeys = keyMap{
	view: chatinput,
	Tab: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "toggle between chat and input"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("⏎", "ask ai"),
	),
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "scroll up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "scroll down"),
	),
	Help: key.NewBinding(
		key.WithKeys("ctrl+h"),
		key.WithHelp("ctrl+h", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
}
