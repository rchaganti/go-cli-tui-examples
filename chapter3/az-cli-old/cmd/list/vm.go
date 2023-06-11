package list

import (
	"az-cli/pkg/az"

	"github.com/spf13/cobra"
)

var listRgCmd = &cobra.Command{
	Use:   "rg",
	Short: "List all Azure resource groups",
	Long:  "List all Azure resource groups",
	Args:  cobra.NoArgs,
	RunE:  az.Info,
}

func init() {
	listCmd.AddCommand(listRgCmd)
}
