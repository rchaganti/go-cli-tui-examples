package cmd

import (
	cli "github.com/urfave/cli/v2"
)

var rootApp = &cli.App{
	Name:        "gai",
	Usage:       "Your command line interface to Google's Gemini AI",
	Description: "This is a command line interface to Google's Gemini AI",
	Flags: []cli.Flag{
		apiKeyFlag,
	},
	Commands: []*cli.Command{
		modelCommand,
		generateCommand,
	},
	EnableBashCompletion: true,
}

func Run(args []string) error {
	return rootApp.Run(args)

}
