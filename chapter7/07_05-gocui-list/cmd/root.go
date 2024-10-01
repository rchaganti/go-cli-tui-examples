package cmd

import (
	"gocui-list/ui"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gocui-list",
	Short: "gocui-list is an example application that demonstrates how to use the gocui library to create a list view",
	Run: func(cmd *cobra.Command, args []string) {
		err := ui.InvokeList()
		if err != nil {
			panic(err)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}
