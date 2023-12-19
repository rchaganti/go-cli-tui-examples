package cmd

import (
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:       "completion [SHELL]",
	Short:     "Prints shell completion scripts",
	Long:      "Prints shell completion scripts",
	ValidArgs: []string{"bash", "powershell"},
	Annotations: map[string]string{
		"commandType": "main",
	},
	Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		switch args[0] {
		case "bash":
			_ = cmd.Root().GenBashCompletion(cmd.OutOrStdout())
		case "powershell":
			_ = cmd.Root().GenPowerShellCompletion(cmd.OutOrStdout())
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
