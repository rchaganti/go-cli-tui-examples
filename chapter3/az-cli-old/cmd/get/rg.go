package get

import (
	"az-cli/pkg/az"

	"github.com/spf13/cobra"
)

var getRgCmd = &cobra.Command{
	Use:   "rg",
	Short: "Get information about an Azure resource group",
	Long:  "Get detailed information about an Azure resource group",
	Args:  cobra.NoArgs,
	RunE:  az.Info,
}

func init() {
	getCmd.AddCommand(getRgCmd)
}
