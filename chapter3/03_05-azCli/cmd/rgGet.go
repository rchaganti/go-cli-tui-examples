package cmd

import (
	"azCli/pkg/az"

	"github.com/spf13/cobra"
)

var rgGetCmd = &cobra.Command{
	Use:   "resourceGroup [OPTIONS...]",
	Short: "Get an Azure resource group information.",
	Long:  "Get an Azure resource group details",
	Run: func(cmd *cobra.Command, args []string) {
		subscriptionId, _ := cmd.Flags().GetString("subscriptionId")
		name, _ := cmd.Flags().GetString("name")
		az.GetResourceGroup(subscriptionId, name)
	},
	Aliases: []string{"rg", "resourceGroups"},
}

func init() {
	getCmd.AddCommand(rgGetCmd)
}
