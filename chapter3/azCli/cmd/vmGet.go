package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// vmGetCmd represents the vmGet command
var vmGetCmd = &cobra.Command{
	Use:   "vm",
	Short: "Get an Azure virtual machine information.",
	Long:  "Get an Azure virtual machine details",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vmGet called")
	},
}

var vmName string

func init() {
	getCmd.AddCommand(vmGetCmd)
	vmGetCmd.Flags().StringVarP(&vmName, "name", "n", "", "Help message for toggle")
}
