package main

import (
	"docker-containermgr/cmd"

	"github.com/docker/cli/cli-plugins/manager"
	"github.com/docker/cli/cli-plugins/plugin"
)

var (
	pluginMetadata = manager.Metadata{
		SchemaVersion:    "0.1.0",
		Version:          "0.1.0",
		Vendor:           "Cloud Native Central",
		ShortDescription: "A TUI tool to manage docker containers",
		URL:              "github.com/rchaganti/docker-containermgr",
	}

	containerCount bool
)

func main() {
	plugin.Run(
		cmd.NewRootCommand,
		pluginMetadata,
	)
}
