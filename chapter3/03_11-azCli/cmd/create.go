package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create [OPTIONS]",
	Short: "Create a specific type of Azure resource.",
	Long:  "Create a specific type of Azure resource.",
	Args:  cobra.ExactArgs(1),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if os.Getenv("AZURE_CLIENT_ID") == "" || os.Getenv("AZURE_CLIENT_SECRET") == "" || os.Getenv("AZURE_TENANT_ID") == "" {
			return fmt.Errorf("One or more of required environment variables are not set. Please set AZURE_CLIENT_ID, AZURE_CLIENT_SECRET, and AZURE_TENANT_ID")
		}
		return nil
	},
	SuggestFor: []string{"new"},
	GroupID:    "create",
}

func init() {
	createCmd.PersistentFlags().StringVarP(&location, "location", "l", "", "Location where the Azure resource group should be created.")
	createCmd.MarkPersistentFlagRequired("location")

	createCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Name of the Azure resource")
	createCmd.MarkPersistentFlagRequired("name")

	rootCmd.AddCommand(createCmd)
}
