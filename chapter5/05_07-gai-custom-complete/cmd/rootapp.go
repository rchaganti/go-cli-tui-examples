package cmd

import (
	"fmt"

	cli "github.com/urfave/cli/v2"
)

var appHelpTemplate = fmt.Sprintf(`%s

WEBSITE: http://awesometown.example.com

SUPPORT: support@awesometown.example.com

`, cli.AppHelpTemplate)

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
	Suggest:              true,
}

func Run(args []string) error {
	return rootApp.Run(args)
}

func init() {
	cli.AppHelpTemplate = appHelpTemplate
}
