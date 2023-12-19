package cmd

import (
	"azCli/pkg/az"

	"github.com/spf13/cobra"
)

var vmListCmd = &cobra.Command{
	Use:   "virtualMachine [OPTIONS...]",
	Short: "List all Azure virtual machines.",
	Long:  "List all Azure virtual machines  in a resource group.",
	Run: func(cmd *cobra.Command, args []string) {
		subscriptionId, _ := cmd.Flags().GetString("subscriptionId")
		resourceGroupName, _ := cmd.Flags().GetString("resourceGroupName")
		az.ListVirtualMachine(subscriptionId, resourceGroupName)
	},
	Aliases: []string{"vm", "virtualMachines"},
}

func init() {
	listCmd.AddCommand(vmListCmd)
	vmListCmd.Flags().StringVarP(&resourceGroupName, "resourceGroupName", "r", "", "Name of the eresource group where the VM is provisioned.")

	vmListCmd.RegisterFlagCompletionFunc("resourceGroupName", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		subscriptionId, _ := cmd.Flags().GetString("subscriptionId")
		return az.ListResourceGroupName(subscriptionId), cobra.ShellCompDirectiveNoFileComp
	})
}
