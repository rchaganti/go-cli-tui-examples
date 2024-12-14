package cmd

import (
	"contextmgr/ui"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/tools/clientcmd"
)

var configFlags = genericclioptions.NewConfigFlags(true)

var rootCmd = &cobra.Command{
	Use:   "contextmgr",
	Short: "A simple CLI tool to manage kubernetes contexts",
	Long:  "A simple CLI tool to manage kubernetes contexts",
	RunE: func(cmd *cobra.Command, args []string) error {
		var c tea.Model
		var err error

		kubeconfig, err := cmd.Flags().GetString("kubeconfig")
		if err != nil {
			return err
		}

		if kubeconfig == "" {
			kubeconfig = configFlags.ToRawKubeConfigLoader().ConfigAccess().GetDefaultFilename()
		}

		config, err := clientcmd.LoadFromFile(kubeconfig)
		if err != nil {
			return err
		}

		m := ui.NewContextModel(kubeconfig, config)

		if c, err = tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}

		if m, ok := c.(*ui.ContextModel); ok && m.Error != nil {
			fmt.Println("Error:", m.Error)
			os.Exit(1)
		}

		return nil
	},
}

func Execute() error {
	err := rootCmd.Execute()
	if err != nil {
		return err
	}

	return nil
}

func init() {
	configFlags.AddFlags(rootCmd.PersistentFlags())
}
