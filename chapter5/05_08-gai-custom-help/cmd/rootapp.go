package cmd

import (
	"fmt"
	"io"

	cli "github.com/urfave/cli/v2"
)

var banner = `
________    _____  .___ 
/  _____/   /  _  \ |   |
/   \  ___  /  /_\  \|   |
\    \_\  \/    |    \   |
\______  /\____|__  /___|
	   \/         \/     
To know more: https://leanpub.com/go-cli-tui
`
var appHelpTemplate = fmt.Sprintf(`%s

%s
`, banner, cli.AppHelpTemplate)

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

	cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
		switch data := data.(type) {
		case *cli.App:
			fmt.Fprintf(w, `%s
%s - %s

`, banner, data.Name, data.Usage)
		case *cli.Command:
			fmt.Fprintf(w, `%s
%s - %s

`, banner, data.Name, data.Usage)
		}
	}
}
