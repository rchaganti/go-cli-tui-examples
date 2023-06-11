package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "az-cli",
	Short: "Azure CLI re-invented",
	Long:  "A modern implementation of Azure CLI written in Go",
	Args:  cobra.NoArgs,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ooops. There was an error while executing command '%s'", err)
		os.Exit(1)
	}
}
