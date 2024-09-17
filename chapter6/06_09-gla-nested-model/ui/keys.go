package ui

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	view   viewState
	Up     key.Binding
	Down   key.Binding
	Enter  key.Binding
	Tab    key.Binding
	Delete key.Binding
	Toggle key.Binding
	Help   key.Binding
	Quit   key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	if k.view == chatview {
		return []key.Binding{
			k.Up, k.Down, k.Tab, k.Help, k.Quit,
		}
	} else if k.view == chatinput {
		return []key.Binding{
			k.Enter, k.Tab, k.Help, k.Quit,
		}
	} else if k.view == listview {
		return []key.Binding{
			k.Up, k.Down, k.Enter, k.Toggle, k.Help, k.Quit,
		}
	}

	return []key.Binding{}
}

func (k keyMap) FullHelp() [][]key.Binding {
	if k.view == chatview {
		return [][]key.Binding{
			{k.Up, k.Down, k.Tab, k.Toggle},
			{k.Help, k.Quit},
		}
	} else if k.view == listview {
		return [][]key.Binding{
			{k.Up, k.Down, k.Enter, k.Delete},
			{k.Toggle, k.Help, k.Quit},
		}
	} else if k.view == chatinput {
		return [][]key.Binding{
			{k.Enter, k.Tab},
			{k.Help, k.Quit},
		}
	}

	return [][]key.Binding{}
}

var chatKeys = keyMap{
	view: chatinput,
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "scroll up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "scroll down"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("⏎", "ask ai"),
	),
	Tab: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "toggle between chat and input"),
	),
	Toggle: key.NewBinding(
		key.WithKeys("ctrl+l"),
		key.WithHelp("ctrl+l", "toggle to session list"),
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

var listKeys = keyMap{
	view: listview,
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "scroll up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "scroll down"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select session"),
	),
	Delete: key.NewBinding(
		key.WithKeys("delete"),
		key.WithHelp("delete", "delete session"),
	),
	Toggle: key.NewBinding(
		key.WithKeys("ctrl+l"),
		key.WithHelp("ctrl+l", "toggle to chat view"),
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
