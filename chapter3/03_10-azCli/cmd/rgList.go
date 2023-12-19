package cmd

import (
	"azCli/pkg/az"

	"github.com/spf13/cobra"
)

var rgListCmd = &cobra.Command{
	Use:   "resourceGroup [OPTIONS...]",
	Short: "List all Azure resource groups",
	Long:  "List all Azure resource groups in a subscription",
	Run: func(cmd *cobra.Command, args []string) {
		subscriptionId, _ := cmd.Flags().GetString("subscriptionId")
		az.ListResourceGroup(subscriptionId)
	},
	Aliases: []string{"rg", "resourceGroups"},
}

func init() {
	listCmd.AddCommand(rgListCmd)
}
