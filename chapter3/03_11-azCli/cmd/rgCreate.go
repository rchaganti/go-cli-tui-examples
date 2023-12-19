package cmd

import (
	"azCli/pkg/az"

	"github.com/spf13/cobra"
)

var rgCreateCmd = &cobra.Command{
	Use:   "resourceGroup [OPTIONS...]",
	Short: "Create an Azure resource group",
	Long:  "Create an Azure resource group.",
	Run: func(cmd *cobra.Command, args []string) {
		subscriptionId, _ := cmd.Flags().GetString("subscriptionId")
		name, _ := cmd.Flags().GetString("name")
		location, _ := cmd.Flags().GetString("location")
		az.CreateResourceGroup(subscriptionId, name, location)
	},
	Aliases: []string{"rg", "resourceGroups"},
}

func init() {
	createCmd.AddCommand(rgCreateCmd)
}
