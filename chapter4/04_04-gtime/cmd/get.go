package cmd

import (
	"fmt"

	"github.com/rchaganti/gtime/pkg/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get configuration parameters",
	Long:  "Get the configuration parameters used by gtime",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := viper.GetString("config-path")

		configDir, configName, configType := helper.ParseConfigPath(configPath)
		viper.SetConfigName(configName)
		viper.AddConfigPath(configDir)
		viper.SetConfigType(configType)

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				fmt.Printf("Configuration file %s not found\n", configPath)
			} else {
				fmt.Printf("Error reading configuration file: %s\n", err)
			}
		} else {
			config := Config{}
			err := viper.Unmarshal(&config)

			if err != nil {
				fmt.Printf("Error unmarshalling configuration file: %s\n", err)
			}
			fmt.Printf("Timezones: %v\n", config.Timezones)
			fmt.Printf("Output: %s\n", config.Output)
		}
	},
}

func init() {
	configCmd.AddCommand(getCmd)
}
