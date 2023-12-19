package cmd

import (
	"azCli/pkg/az"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var vmGetCmd = &cobra.Command{
	Use:   "virtualMachine [OPTIONS...]",
	Short: "Get an Azure virtual machine information.",
	Long:  "Get an Azure virtual machine details",
	Run: func(cmd *cobra.Command, args []string) {
		subscriptionId, _ := cmd.Flags().GetString("subscriptionId")
		var name string
		if len(args) == 0 {
			log.Fatal("Virtual Machine name is required")
		} else {
			name = args[0]
		}

		resourceGroupName, _ := cmd.Flags().GetString("resourceGroupName")
		az.GetVirtualMachine(subscriptionId, resourceGroupName, name)
	},
	Aliases: []string{"vm", "virtualMachines"},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		var comps []string
		subscriptionId, _ := cmd.Flags().GetString("subscriptionId")
		resourceGroupName, _ := cmd.Flags().GetString("resourceGroupName")
		vms := az.ListVirtualMachineName(subscriptionId, resourceGroupName)
		for _, vm := range vms {
			if toComplete == "" {
				comps = append(comps, vm)
			} else if strings.HasPrefix(vm, toComplete) {
				comps = append(comps, vm)
			}
		}

		return comps, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	getCmd.AddCommand(vmGetCmd)
	vmGetCmd.Flags().StringVarP(&resourceGroupName, "resourceGroupName", "r", "", "Name of the resource group where the VM is provisioned.")

	vmGetCmd.RegisterFlagCompletionFunc("resourceGroupName", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		subscriptionId, _ := cmd.Flags().GetString("subscriptionId")
		return az.ListResourceGroupName(subscriptionId), cobra.ShellCompDirectiveNoFileComp
	})
}
