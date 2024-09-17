package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	cli "github.com/urfave/cli/v2"
)

var name string

func main() {
	app := &cli.App{
		Name:        "hello",
		Description: "Greets a Gopher",
		Usage:       "Shows a greeting message",
		UsageText:   "hello [command options]",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "name",
				Aliases:     []string{"n"},
				Required:    true,
				Destination: &name,
			},
		},
		Action: func(c *cli.Context) error {
			if strings.TrimSpace(name) != "" {
				fmt.Printf("Hello, %s\n", name)
			} else {
				fmt.Printf("%s cannot be empty\n\n", "--name")
				cli.ShowAppHelpAndExit(c, 1)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
