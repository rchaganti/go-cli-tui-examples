package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// vmListCmd represents the vmList command
var vmListCmd = &cobra.Command{
	Use:   "vm",
	Short: "List all Azure virtual machines.",
	Long:  "List all Azure virtual machines  in a resource group.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vmList called")
	},
}

func init() {
	listCmd.AddCommand(vmListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// vmListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// vmListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
