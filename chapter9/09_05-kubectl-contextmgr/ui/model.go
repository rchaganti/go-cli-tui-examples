package ui

import (
	helper "contextmgr/internal"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	"k8s.io/client-go/tools/clientcmd/api"
)

type ContextModel struct {
	Table      table.Model
	Contexts   []helper.Context
	Config     *api.Config
	Kubeconfig string
	Error      error
	Help       help.Model
	Keys       keyMap
	Spinner    spinner.Model
	Refreshing bool
}

var columns = []table.Column{
	{Title: "Status", Width: 7},
	{Title: "Context", Width: 15},
	{Title: "Cluster", Width: 15},
}

func (c ContextModel) Init() tea.Cmd {
	return nil
}

func (c ContextModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			cmds = append(cmds, tea.Quit)
		case "enter":
			selectedContext := c.Table.SelectedRow()[1]
			if selectedContext != c.Config.CurrentContext {
				cmds = append(cmds, c.switchContext())
			} else {
				return c, nil
			}
		case "r":
			c.Refreshing = true
			c.Spinner, cmd = c.Spinner.Update(msg)
			cmds = append(cmds, cmd, c.Spinner.Tick, c.refreshContexts())
		default:
			c.Table, cmd = c.Table.Update(msg)
			cmds = append(cmds, cmd)
		}

	case ContextSwitchedMsg:
		c.Config.CurrentContext = msg.currentContext
		c.updateTableRows()

	case ContextRefreshedMsg:
		c.Refreshing = false
		c.Contexts = msg.contexts
		c.Config = msg.config
		c.updateTableRows()

	case ErrMsg:
		c.Error = msg.Err
		return c, tea.Quit

	case spinner.TickMsg:
		c.Spinner, cmd = c.Spinner.Update(msg)
		cmds = append(cmds, cmd)

	default:
		c.Table, cmd = c.Table.Update(msg)
		cmds = append(cmds, cmd)
	}
	return c, tea.Batch(cmds...)
}

func (c ContextModel) View() string {
	tableView := "\n" + c.Table.View()
	if c.Refreshing {
		return lipgloss.JoinVertical(lipgloss.Left,
			tableView,
			c.helpView(),
			c.spinnerView(),
		)
	} else {
		return lipgloss.JoinVertical(lipgloss.Left,
			tableView,
			c.helpView(),
		)
	}
}

func NewContextModel(kubeconfig string, config *api.Config) *ContextModel {
	contexts := helper.GetContext(config)
	rows := createTableRows(contexts)

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	t.SetStyles(newTableStyle())

	help := help.New()
	help.Styles.FullKey = helpKeysStyle
	help.Styles.ShortKey = helpKeysStyle
	help.Styles.ShortDesc = helpDescStyle
	help.Styles.FullDesc = helpDescStyle
	help.ShowAll = true

	s := spinner.New()
	s.Spinner = spinner.Ellipsis
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return &ContextModel{
		Table:      t,
		Contexts:   contexts,
		Config:     config,
		Kubeconfig: kubeconfig,
		Help:       help,
		Keys:       listKeys,
		Spinner:    s,
		Refreshing: false,
	}
}
