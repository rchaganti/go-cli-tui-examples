package cmd

import (
	"fmt"

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
	Long:  "gtime - track time across timezones",
	Run: func(cmd *cobra.Command, args []string) {
		target := viper.GetString("target")
		timezones := viper.GetStringSlice("timezones")
		output := viper.GetString("output")

		fmt.Printf("target: %s\ntimezones: %s\noutput: %s\n", target, timezones, output)
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
