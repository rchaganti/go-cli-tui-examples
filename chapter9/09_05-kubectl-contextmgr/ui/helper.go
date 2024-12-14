package ui

import (
	helper "contextmgr/internal"
	"fmt"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"k8s.io/client-go/tools/clientcmd/api"
)

type ContextSwitchedMsg struct {
	currentContext string
}

type ContextRefreshedMsg struct {
	contexts []helper.Context
	config   *api.Config
}

type ErrMsg struct{ Err error }

func (c *ContextModel) switchContext() tea.Cmd {
	return func() tea.Msg {
		selectedContext := c.Table.SelectedRow()[1]
		_, err := helper.SwitchContext(selectedContext, c.Config)
		if err != nil {
			return ErrMsg{Err: err}
		}
		for i, ctx := range c.Contexts {
			if ctx.Name == selectedContext {
				c.Contexts[i].IsCurrent = true
			} else {
				c.Contexts[i].IsCurrent = false
			}
		}
		return ContextSwitchedMsg{
			currentContext: selectedContext,
		}
	}
}

func (c *ContextModel) refreshContexts() tea.Cmd {
	return func() tea.Msg {
		contexts := helper.GetContext(c.Config)
		return ContextRefreshedMsg{contexts: contexts, config: c.Config}
	}
}

func createTableRows(contexts []helper.Context) []table.Row {
	rows := make([]table.Row, len(contexts))

	for i, ctx := range contexts {
		statusIcon := "✖"
		if ctx.Status {
			statusIcon = "✔"
		}

		if ctx.IsCurrent {
			statusIcon += " " + "*"
		}

		rows[i] = table.Row{
			statusIcon,
			ctx.Name,
			ctx.Cluster,
		}
	}
	return rows
}

func (c *ContextModel) updateTableRows() {
	rows := createTableRows(c.Contexts)
	c.Table.SetRows(rows)
}

func (c ContextModel) helpView() string {
	view := c.Help.View(c.Keys)
	return helpStyle.Render(view)
}

func (c ContextModel) spinnerView() string {
	view := fmt.Sprintf("Refreshing contexts%s", c.Spinner.View())
	return view
}
