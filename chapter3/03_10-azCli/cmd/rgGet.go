package cmd

import (
	"azCli/pkg/az"
	"log"

	"github.com/spf13/cobra"
)

var rgGetCmd = &cobra.Command{
	Use:   "resourceGroup [OPTIONS...] RESOURCEGROUPNAME",
	Short: "Get an Azure resource group information.",
	Long:  "Get an Azure resource group details",
	Run: func(cmd *cobra.Command, args []string) {
		subscriptionId, _ := cmd.Flags().GetString("subscriptionId")

		var name string
		if len(args) == 0 {
			log.Fatal("Resource group name is required")
		} else {
			name = args[0]
		}

		az.GetResourceGroup(subscriptionId, name)
	},
	Aliases: []string{"rg", "resourceGroups"},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		subscriptionId, _ := cmd.Flags().GetString("subscriptionId")
		var comps []string
		comps = az.ListResourceGroupName(subscriptionId)

		if len(args) == 0 {
			comps = cobra.AppendActiveHelp(comps, "Specify a resource group name")
		} else if len(args) == 1 {
			comps = cobra.AppendActiveHelp(comps, "This command does not take any more arguments")
		}
		return comps, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	getCmd.AddCommand(rgGetCmd)
}
