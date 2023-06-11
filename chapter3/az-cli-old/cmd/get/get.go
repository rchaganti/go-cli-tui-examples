package get

import (
	cmd "az-cli/cmd"
	"az-cli/pkg/az"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get information about an Azure resource",
	Long:  "Get detailed information about an Azure resource of specifc type",
	Args:  cobra.NoArgs,
	RunE:  az.Info,
}

func init() {
	cmd.RootCmd.AddCommand(getCmd)
}
