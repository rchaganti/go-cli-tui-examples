package cmd

import (
	"gocui-md-reader-list/ui"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gocui-md-reader",
	Short: "A simple markdown reader",
	Run: func(cmd *cobra.Command, args []string) {
		err := ui.NewMDReader()
		if err != nil {
			panic(err)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}
