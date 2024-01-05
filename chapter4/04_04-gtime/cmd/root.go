package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/rchaganti/gtime/pkg/gtime"
	"github.com/rchaganti/gtime/pkg/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Timezones []string `json:"timezones"`
	Output    string   `json:"output"`
}

var (
	target     string
	configPath string
)

var rootCmd = &cobra.Command{
	Use:   "gtime",
	Short: "gtime - track time across timezones",
	Long:  "gtime - track time across timezones",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := viper.GetString("config-path")
		configDir, configName, configType := helper.ParseConfigPath(configPath)

		viper.SetConfigName(configName)
		viper.AddConfigPath(configDir)
		viper.SetConfigType(configType)

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				fmt.Printf("Error reading configuration file: %s\n", err)
			}
		}

		target, _ := cmd.Flags().GetString("target")

		if target == "" {
			target = time.Now().Format("2/1/2006 15:04")
		}

		timezones := viper.GetStringSlice("timezones")
		output := viper.GetString("output")

		err := gtime.ConvertTime(target, timezones, output)
		if err != nil {
			fmt.Printf("Error converting target time to specified timezones: %s", err)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	conf := Config{}

	rootCmd.Flags().StringVarP(&target, "target", "t", "", "Target time to convert to different timezones.")
	rootCmd.PersistentFlags().StringSliceVarP(&conf.Timezones, "timezones", "z", []string{}, "Timezones to convert the target time to.")
	rootCmd.PersistentFlags().StringVarP(&conf.Output, "output", "o", "table", "Output format to use for displaying the time information.")

	rootCmd.PersistentFlags().StringVarP(&configPath, "config-path", "c", "./config.json", "Configuration path to use. Include the file name and extension.")

	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
	viper.BindPFlag("timezones", rootCmd.PersistentFlags().Lookup("timezones"))
	viper.BindPFlag("config-path", rootCmd.PersistentFlags().Lookup("config-path"))

	viper.SetEnvPrefix("GTIME")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()
}
