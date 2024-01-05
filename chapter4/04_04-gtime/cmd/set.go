package cmd

import (
	"fmt"
	"os"

	"github.com/rchaganti/gtime/pkg/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set configuration parameters",
	Long:  "Set the configuration parameters used by gtime",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := viper.GetString("config-path")
		configDir, configName, configType := helper.ParseConfigPath(configPath)

		viper.AddConfigPath(configDir)
		viper.SetConfigName(configName)
		viper.SetConfigType(configType)

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				fmt.Printf("Configuration file %s not found. It will be created\n", configPath)
				if _, err := os.Stat(configDir); os.IsNotExist(err) {
					fmt.Println("creating configuration directory")
					err := os.MkdirAll(configDir, 0755)
					if err != nil {
						fmt.Printf("Error creating configuration directory: %s\n", err)
					}
				}
				err = viper.WriteConfigAs(configPath)
				if err != nil {
					fmt.Printf("Error writing configuration file: %s\n", err)
				}
			} else {
				fmt.Printf("Error reading configuration file: %s\n", err)
			}
		} else {
			fmt.Printf("Configuration file %s will be updated\n", configPath)
			err = viper.WriteConfig()
			if err != nil {
				fmt.Printf("Error writing configuration file: %s\n", err)
			}
		}
	},
}

func init() {
	configCmd.AddCommand(setCmd)
}
