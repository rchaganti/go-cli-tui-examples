package cmd

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get information about an Azure resource",
	Long:  "Get detailed information about a resource from an Azure resource group",
	Args:  cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.PersistentFlags().StringVarP(&subscriptionId, "subscriptionId", "s", "", "ID of the Azure subscription where the Azure resources are provisioned.")
	getCmd.MarkPersistentFlagRequired("subscriptionId")

	getCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Name of the Azure resource")
	getCmd.MarkPersistentFlagRequired("name")
}
