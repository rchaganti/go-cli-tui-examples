package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	cli "github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "hello"
	app.Description = "Greets a human or the world"
	app.Usage = "Shows a greeting message"
	app.UsageText = "hello [arguments]"
	app.Action = func(c *cli.Context) error {
		var name string
		if c.Args().Present() {
			name = c.Args().Get(0)
			if strings.Trim(name, " ") == "" {
				name = "World"
			}
			fmt.Printf("Hello, %s\n", name)
			return nil
		} else {
			cli.ShowAppHelpAndExit(c, 1)
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
