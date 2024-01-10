package cmd

import (
	cli "github.com/urfave/cli/v2"
)

var (
	apiKeyFlag = &cli.StringFlag{
		Name:    "api-key",
		Aliases: []string{"a"},
		Usage:   "API key for authenticating to Google Gemini AI",
		EnvVars: []string{"GAI_API_KEY"},
	}

	outputFlag = &cli.StringFlag{
		Name:    "output",
		Aliases: []string{"o"},
		Usage:   "Output format for displaying model information",
		Value:   "json",
	}

	promptFlag = &cli.StringFlag{
		Name:    "prompt",
		Aliases: []string{"p"},
		Usage:   "Prompt for generating response from Google Gemini AI",
	}

	promptFromFileFlag = &cli.StringFlag{
		Name:    "prompt-from-file",
		Aliases: []string{"f"},
		Usage:   "Text file prompt for generating response from Google Gemini AI",
	}

	imageForPromptFlag = &cli.StringFlag{
		Name:    "image-for-prompt",
		Aliases: []string{"i"},
		Usage:   "Image file as input to the prompt for generating response from Google Gemini AI",
	}
)
