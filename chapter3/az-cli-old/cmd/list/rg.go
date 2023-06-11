package list

import (
	"az-cli/pkg/az"

	"github.com/spf13/cobra"
)

var listVmCmd = &cobra.Command{
	Use:   "vm",
	Short: "List all Azure virtual machines",
	Long:  "List all Azure virtual machines in a resource group",
	Args:  cobra.NoArgs,
	RunE:  az.Info,
}

func init() {
	listCmd.AddCommand(listVmCmd)
}
