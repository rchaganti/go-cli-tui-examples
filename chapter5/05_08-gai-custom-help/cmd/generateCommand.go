package cmd

import (
	"gai/pkg/gai"
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
)

var commandHelpTemplate = `
NAME: {{ .HelpName }}

USAGE: {{template "usageTemplate" .}}
{{if .Category}}
CATEGORY:
{{.Category}}{{end}}
{{if .VisibleFlags}}
OPTIONS:{{template "visibleFlagTemplate" .}}
{{end}}
`

var (
	generateCommand = &cli.Command{
		Name:        "generate",
		Usage:       "Generate response from Google Gemini AI",
		Description: "Generate response from Google Gemini AI",
		Category:    "generate",
		Flags: []cli.Flag{
			promptFlag,
			promptFromFileFlag,
			imageForPromptFlag,
		},
		Action: func(c *cli.Context) error {
			if c.IsSet("prompt-from-file") && c.IsSet("prompt") {
				return cli.Exit("prompt and prompt-from-file are mutually exclusive", 1)
			}

			apiKey := c.String("api-key")
			prompt := c.String("prompt")

			if c.IsSet("prompt-from-file") {
				filePath := c.String("prompt-from-file")
				c, err := os.ReadFile(filePath)
				if err != nil {
					log.Fatal(err)
				}
				prompt = string(c)
			}

			if c.IsSet("image-for-prompt") {
				imagePath := c.String("image-for-prompt")
				resp, err := gai.GenerateFromImage(apiKey, imagePath, prompt)
				if err != nil {
					log.Fatal(err)
				}
				gai.PrintResponse(resp)
			} else {
				resp, err := gai.GenerateFromText(apiKey, prompt)
				if err != nil {
					log.Fatal(err)
				}
				gai.PrintResponse(resp)
			}

			return nil
		},
	}
)

func init() {
	cli.CommandHelpTemplate = commandHelpTemplate
}
