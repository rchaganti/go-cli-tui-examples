package cmd

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get information about an Azure resource",
	Long:  "Get detailed information about a resource from an Azure resource group",
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Name of the Azure resource")
	getCmd.MarkPersistentFlagRequired("name")
}
