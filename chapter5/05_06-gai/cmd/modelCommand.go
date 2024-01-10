package cmd

import (
	"encoding/json"
	"fmt"

	"gai/pkg/gai"

	cli "github.com/urfave/cli/v2"
	yaml "gopkg.in/yaml.v2"
)

var (
	modelCommand = &cli.Command{
		Name:        "model",
		Usage:       "Manage [get | list] Google Gemini AI models",
		Description: "Manage [get | list] Google Gemini AI models",
		Category:    "model",
		Flags:       []cli.Flag{},
		Subcommands: []*cli.Command{
			modelGetCommand,
			modelListCommand,
		},
	}

	modelGetCommand = &cli.Command{
		Name:        "get",
		Usage:       "Get a Google Gemini AI model",
		Description: "Get a Google Gemini AI model",
		Category:    "model",
		Flags: []cli.Flag{
			modelNameFlag,
			outputFlag,
		},
		Action: func(c *cli.Context) error {
			var modelinfo []byte

			apiKey := c.String("api-key")
			output := c.String("output")
			modelName := c.String("model-name")
			model := gai.GetModel(apiKey, modelName)

			if output == "json" {
				modelinfo, _ = json.MarshalIndent(model, "", "  ")
			} else if output == "yaml" {
				modelinfo, _ = yaml.Marshal(model)
			} else {
				return fmt.Errorf("invalid output format: %s", output)
			}
			fmt.Println(string(modelinfo))
			return nil
		},
	}

	modelListCommand = &cli.Command{
		Name:        "list",
		Usage:       "List Google Gemini AI models",
		Description: "List Google Gemini AI models",
		Category:    "model",
		Flags: []cli.Flag{
			outputFlag,
		},
		Action: func(c *cli.Context) error {
			var modelList []byte
			apiKey := c.String("api-key")
			output := c.String("output")
			models := gai.ListModels(apiKey)
			if output == "json" {
				modelList, _ = json.MarshalIndent(models, "", "  ")
			} else if output == "yaml" {
				modelList, _ = yaml.Marshal(models)
			} else {
				return fmt.Errorf("invalid output format: %s", output)
			}

			fmt.Println(string(modelList))
			return nil
		},
	}
)
