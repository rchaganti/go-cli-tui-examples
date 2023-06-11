package get

import (
	"az-cli/pkg/az"

	"github.com/spf13/cobra"
)

var getVmCmd = &cobra.Command{
	Use:   "vm",
	Short: "Get information about an Azure virtual machine",
	Long:  "Get detailed information about an Azure virtual machine",
	Args:  cobra.NoArgs,
	RunE:  az.Info,
}

func init() {
	getCmd.AddCommand(getVmCmd)
}
