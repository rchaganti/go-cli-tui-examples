package cmd

import (
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:        "delete",
	Short:      "Delete a specific type of Azure resource.",
	Long:       "Delete a specific type of Azure resource.",
	Args:       cobra.ExactArgs(1),
	SuggestFor: []string{"remove"},
	GroupID:    "delete",
}

func init() {
	deleteCmd.PersistentFlags().StringVarP(&subscriptionId, "subscriptionId", "s", "", "ID of the Azure subscription where the Azure resources are provisioned.")
	deleteCmd.MarkPersistentFlagRequired("subscriptionId")

	deleteCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Name of the Azure resource")
	deleteCmd.MarkPersistentFlagRequired("name")

	rootCmd.AddCommand(deleteCmd)
}
