package cmd

import (
	"azCli/pkg/az"

	"github.com/spf13/cobra"
)

var vmGetCmd = &cobra.Command{
	Use:   "virtualMachine [OPTIONS...]",
	Short: "Get an Azure virtual machine information.",
	Long:  "Get an Azure virtual machine details",
	Run: func(cmd *cobra.Command, args []string) {
		subscriptionId, _ := cmd.Flags().GetString("subscriptionId")
		name, _ := cmd.Flags().GetString("name")
		resourceGroupName, _ := cmd.Flags().GetString("resourceGroupName")
		az.GetVirtualMachine(subscriptionId, resourceGroupName, name)
	},
	Aliases: []string{"vm", "virtualMachines"},
}

func init() {
	getCmd.AddCommand(vmGetCmd)
	vmGetCmd.Flags().StringVarP(&resourceGroupName, "resourceGroupName", "r", "", "Name of the resource group where the VM is provisioned.")
	vmGetCmd.MarkFlagRequired("resourceGroupName")
}
