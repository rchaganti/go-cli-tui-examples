package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	subscriptionId    string
	resourceGroupName string
	name              string
)

var rootCmd = &cobra.Command{
	Use:   "az",
	Short: "A re-imagined Azure CLI",
	Long:  "Azure CLI written in Go language using the Cobra package.",
	Args:  cobra.NoArgs,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&subscriptionId, "subscriptionId", "s", "", "ID of the Azure subscription where the Azure resources are provisioned.")
	rootCmd.MarkPersistentFlagRequired("subscriptionId")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
