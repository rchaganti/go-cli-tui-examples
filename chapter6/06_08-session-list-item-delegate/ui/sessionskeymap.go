package ui

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Up     key.Binding
	Down   key.Binding
	Enter  key.Binding
	Delete key.Binding
	Left   key.Binding
	Right  key.Binding
	Help   key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Up, k.Down, k.Enter, k.Help,
	}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Enter, k.Delete},
		{k.Left, k.Right, k.Help},
	}
}

var sessionKeys = KeyMap{
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
		key.WithHelp("enter", "select"),
	),
	Delete: key.NewBinding(
		key.WithKeys("delete"),
		key.WithHelp("delete", "delete session from history"),
	),
	Left: key.NewBinding(
		key.WithKeys("left"),
		key.WithHelp("←", "previous page"),
	),
	Right: key.NewBinding(
		key.WithKeys("right"),
		key.WithHelp("→", "next page"),
	),
	Help: key.NewBinding(
		key.WithKeys("ctrl+h"),
		key.WithHelp("ctrl+h", "toggle help"),
	),
}
