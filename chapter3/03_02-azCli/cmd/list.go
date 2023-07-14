package cmd

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List a specific type of Azure resources",
	Long:  "List a specific type of resources from an Azure resource group.",
}

func init() {
	rootCmd.AddCommand(listCmd)
}
