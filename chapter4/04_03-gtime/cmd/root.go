package cmd

import (
	"fmt"
	"time"

	"github.com/rchaganti/gtime/pkg/gtime"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	target    string
	timezones []string
	output    string
)

const (
	configPath = "$HOME/.gtime"
	configName = "config"
)

var rootCmd = &cobra.Command{
	Use:   "gtime",
	Short: "gtime - track time across timezones",
	Long:  "gtime - track time across timezones to help you schedule meetings",
	Run: func(cmd *cobra.Command, args []string) {
		target := viper.GetString("target")

		if target == "" {
			target = time.Now().Format("2/1/2006 15:04")
		}
		timezones := viper.GetStringSlice("timezones")
		output := viper.GetString("output")

		err := gtime.ConvertTime(target, timezones, output)
		if err != nil {
			fmt.Printf("Error converting taget time to specified timezones: %s", err)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	rootCmd.Flags().StringVarP(&target, "target", "t", "", "Target time to convert to different timezones.")
	rootCmd.PersistentFlags().StringSliceVarP(&timezones, "timezones", "z", []string{}, "Timezones to convert the target time to.")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "table", "Output format to use for displaying the time information.")

	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
	viper.BindPFlag("timezones", rootCmd.PersistentFlags().Lookup("timezones"))
	viper.BindPFlag("target", rootCmd.Flags().Lookup("target"))

	viper.SetEnvPrefix("GTIME")
	viper.AutomaticEnv()
}
