package cmd

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List a specific type of Azure resources",
	Long:  "List a specific type of resources from an Azure resource group.",
	Args:  cobra.ExactArgs(1),
}

func init() {
	listCmd.PersistentFlags().StringVarP(&subscriptionId, "subscriptionId", "s", "", "ID of the Azure subscription where the Azure resources are provisioned.")
	listCmd.MarkPersistentFlagRequired("subscriptionId")

	rootCmd.AddCommand(listCmd)
}
