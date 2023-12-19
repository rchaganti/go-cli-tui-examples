package cmd

import (
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:        "delete",
	Short:      "Delete a specific type of Azure resource.",
	Long:       "Delete a specific type of Azure resource.",
	Args:       cobra.ExactArgs(1),
	SuggestFor: []string{"remove"},
	GroupID:    "delete",
}

func init() {
	deleteCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Name of the Azure resource")

	deleteCmd.PersistentFlags().BoolVarP(&all, "all", "a", false, "Delete all resource groups in the subscription.")

	deleteCmd.MarkFlagsMutuallyExclusive("name", "all")
	deleteCmd.MarkFlagsOneRequired("name", "all")

	rootCmd.AddCommand(deleteCmd)
}
