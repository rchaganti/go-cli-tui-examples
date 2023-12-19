package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [OPTIONS] [SUBCOMMAND]",
	Short: "List a specific type of Azure resource.",
	Long:  "List a specific type of Azure resource.",
	Args:  cobra.ExactArgs(1),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if os.Getenv("AZURE_CLIENT_ID") == "" || os.Getenv("AZURE_CLIENT_SECRET") == "" || os.Getenv("AZURE_TENANT_ID") == "" {
			return fmt.Errorf("One or more of required environment variables are not set. Please set AZURE_CLIENT_ID, AZURE_CLIENT_SECRET, and AZURE_TENANT_ID")
		}
		return nil
	},
	GroupID: "read",
}

func init() {
	rootCmd.AddCommand(listCmd)
}
