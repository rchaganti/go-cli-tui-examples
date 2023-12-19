package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [OPTIONS]",
	Short: "Delete a specific type of Azure resource.",
	Long:  "Delete a specific type of Azure resource.",
	Args:  cobra.ExactArgs(1),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if os.Getenv("AZURE_CLIENT_ID") == "" || os.Getenv("AZURE_CLIENT_SECRET") == "" || os.Getenv("AZURE_TENANT_ID") == "" {
			return fmt.Errorf("One or more of required environment variables are not set. Please set AZURE_CLIENT_ID, AZURE_CLIENT_SECRET, and AZURE_TENANT_ID")
		}
		return nil
	},
	SuggestFor: []string{"remove"},
	GroupID:    "delete",
}

func init() {
	deleteCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Name of the Azure resource")

	deleteCmd.PersistentFlags().BoolVarP(&all, "all", "a", false, "Delete all resource groups in the subscription.")

	deleteCmd.MarkFlagsMutuallyExclusive("name", "all")
	deleteCmd.MarkFlagsOneRequired("name", "all")

	rootCmd.AddCommand(deleteCmd)
}
