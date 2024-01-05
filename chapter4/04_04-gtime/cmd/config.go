package cmd

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage the configuration",
	Long:  "Manage (get/set/delete) the configuration parameters used by gtime",
}

func init() {
	rootCmd.AddCommand(configCmd)
}
