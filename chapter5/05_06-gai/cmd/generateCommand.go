package cmd

import (
	"gai/pkg/gai"
	"os"

	cli "github.com/urfave/cli/v2"
)

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
					panic(err)
				}
				prompt = string(c)
			}

			if c.IsSet("image-for-prompt") {
				imagePath := c.String("image-for-prompt")
				resp := gai.GenerateFromImage(apiKey, imagePath, prompt)
				gai.PrintResponse(resp)
			} else {
				resp := gai.GenerateFromText(apiKey, prompt)
				gai.PrintResponse(resp)
			}

			return nil
		},
	}
)
