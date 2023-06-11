package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var resourceGroup string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "azCli",
	Short: "A re-imagined Azure CLI",
	Long:  "Azure CLI written in Go language using the Cobra package.",
	Args:  cobra.NoArgs,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&resourceGroup, "resourceGroup", "r", "", "Name of the resource group where the Azure VMs are provisioned.")
}
