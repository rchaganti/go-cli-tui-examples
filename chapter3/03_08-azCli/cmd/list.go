package cmd

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List a specific type of Azure resource.",
	Long:    "List a specific type of Azure resource.",
	Args:    cobra.ExactArgs(1),
	GroupID: "read",
}

func init() {
	rootCmd.AddCommand(listCmd)
}
