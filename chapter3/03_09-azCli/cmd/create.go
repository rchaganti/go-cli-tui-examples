package cmd

import (
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:        "create [OPTIONS]",
	Short:      "Create a specific type of Azure resource.",
	Long:       "Create a specific type of Azure resource.",
	Args:       cobra.ExactArgs(1),
	SuggestFor: []string{"new"},
	GroupID:    "create",
}

func init() {
	createCmd.PersistentFlags().StringVarP(&location, "location", "l", "", "Location where the Azure resource group should be created.")
	createCmd.MarkPersistentFlagRequired("location")

	createCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Name of the Azure resource")
	createCmd.MarkPersistentFlagRequired("name")

	rootCmd.AddCommand(createCmd)
}
