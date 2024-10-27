package cmd

import (
	containers "docker-containermgr/internal"
	"docker-containermgr/ui"

	"github.com/docker/cli/cli/command"
	"github.com/spf13/cobra"

	tea "github.com/charmbracelet/bubbletea"
)

func NewRootCommand(dockerCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "containermgr",
		Short: "A TUI tool to manage docker containers",
		Long:  "containermgr is a TUI tool to manage docker containers running locally or remotely.",
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := containers.GetContainers(dockerCli)
			if err != nil {
				return err
			}

			m := ui.NewContainerMgr(c, dockerCli)
			if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
				return err
			}
			return nil
		},
	}

	return cmd
}
