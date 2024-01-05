package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete gtime configuration file",
	Long:  "Delete gtime configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := viper.GetString("config-path")

		if _, err := os.Stat(configPath); err == nil {
			var resp string
			fmt.Printf("Are you sure you want to delete the configuration file %s? (y/n): ", configPath)
			fmt.Scan(&resp)
			if resp == "y" {
				err := os.Remove(configPath)
				if err != nil {
					fmt.Printf("Error deleting configuration file: %s\n", err)
				} else {
					fmt.Printf("%s deleted successfully\n", configPath)
				}
			}
		} else {
			fmt.Println("No configuration file to delete")
		}
	},
}

func init() {
	configCmd.AddCommand(deleteCmd)
}
