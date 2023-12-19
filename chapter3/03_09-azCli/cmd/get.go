package cmd

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:     "get [OPTIONS] [SUBCOMMAND]",
	Short:   "Get information about an Azure resource.",
	Long:    "Get detailed information about a resource.",
	Args:    cobra.ExactArgs(1),
	GroupID: "read",
}

func init() {
	rootCmd.AddCommand(getCmd)
}
