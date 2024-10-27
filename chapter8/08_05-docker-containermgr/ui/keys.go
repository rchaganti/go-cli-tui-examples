package ui

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	view  state
	Up    key.Binding
	Down  key.Binding
	Left  key.Binding
	Right key.Binding
	Enter key.Binding
	Stop  key.Binding
	Tab   key.Binding
	Help  key.Binding
	Quit  key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	if k.view == listView {
		return []key.Binding{
			k.Up, k.Down, k.Tab, k.Help, k.Quit,
		}
	} else if k.view == tabView {
		return []key.Binding{
			k.Up, k.Down, k.Left, k.Right, k.Tab, k.Help, k.Quit,
		}
	}

	return []key.Binding{}
}

func (k keyMap) FullHelp() [][]key.Binding {
	if k.view == listView {
		return [][]key.Binding{
			{k.Up, k.Down, k.Tab, k.Enter},
			{k.Stop, k.Help, k.Quit},
		}
	} else if k.view == tabView {
		return [][]key.Binding{
			{k.Up, k.Down, k.Left, k.Right},
			{k.Tab, k.Help, k.Quit},
		}
	}

	return [][]key.Binding{}
}

var listKeys = keyMap{
	view: listView,
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
		key.WithHelp("⏎", "select container"),
	),
	Stop: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "stop container"),
	),
	Tab: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "toggle to viewer"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
}

var viewerKeys = keyMap{
	view: tabView,
	Up: key.NewBinding(
		key.WithKeys("up", "pgup"),
		key.WithHelp("↑", "scroll up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "pgdown"),
		key.WithHelp("↓", "scroll down"),
	),
	Left: key.NewBinding(
		key.WithKeys("left"),
		key.WithHelp("←", "change tab"),
	),
	Right: key.NewBinding(
		key.WithKeys("right"),
		key.WithHelp("→", "change tab"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
}
