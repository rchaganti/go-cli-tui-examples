package list

import (
	cmd "az-cli/cmd"
	"az-cli/pkg/az"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all Azure resources",
	Long:  "List Azure resources of specifc type",
	Args:  cobra.NoArgs,
	RunE:  az.Info,
}

func init() {
	cmd.RootCmd.AddCommand(listCmd)
}
