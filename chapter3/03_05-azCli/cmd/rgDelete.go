package cmd

import (
	"azCli/pkg/az"

	"github.com/spf13/cobra"
)

var rgDeleteCmd = &cobra.Command{
	Use:   "resourceGroup [OPTIONS...]",
	Short: "Delete a Azure resource group.",
	Long:  "Delete a Azure resource group.",
	Run: func(cmd *cobra.Command, args []string) {
		subscriptionId, _ := cmd.Flags().GetString("subscriptionId")
		name, _ := cmd.Flags().GetString("name")
		deleteAll, _ := cmd.Flags().GetBool("all")
		az.DeleteResourceGroup(subscriptionId, name, deleteAll)
	},
	Aliases: []string{"rg", "resourceGroups"},
}

func init() {
	deleteCmd.AddCommand(rgDeleteCmd)
}
