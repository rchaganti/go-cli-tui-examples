package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/docker/cli/cli-plugins/manager"
	"github.com/docker/cli/cli-plugins/plugin"
	"github.com/docker/cli/cli/command"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

const (
	MB = 1 * 1024 * 1024
)

var (
	pluginMetadata = manager.Metadata{
		SchemaVersion:    "0.1.0",
		Version:          "0.1.0",
		Vendor:           "Cloud Native Central",
		ShortDescription: "Provides a list of container images on the host",
		URL:              "github.com/rchaganti/docker-imagelist",
	}

	containerCount bool
)

func main() {
	plugin.Run(
		newRootCommand,
		pluginMetadata,
	)
}

func newRootCommand(dockerCli command.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "imagelist",
		Short: "Provides a list of container images on the host",
		RunE: func(cmd *cobra.Command, args []string) error {
			return getImageList(dockerCli, cmd)
		},
	}

	cmd.Flags().BoolVarP(&containerCount, "container-count", "c", false, "get the number of [running|stopped|paused|created] containers using an image")
	return cmd
}

func getImageList(dockerCli command.Cli, cmd *cobra.Command) error {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug|tabwriter.TabIndent)

	iscc, _ := cmd.Flags().GetBool("container-count")

	apiClient := dockerCli.Client()
	defer apiClient.Close()

	ctx := context.Background()

	images, err := apiClient.ImageList(context.Background(), image.ListOptions{})
	if err != nil {
		panic(err)
	}

	if iscc {
		fmt.Fprintln(w, "ID", "\t", "Tag", "\t", "Size(MB)", "\t", "RunningContainers", "\t", "StoppedContainers", "\t", "PausedContainers", "\t", "CreatedContainers")
	} else {
		fmt.Fprintln(w, "ID", "\t", "Tag", "\t", "Size(MB)")
	}

	for _, img := range images {
		var scc, rcc, pcc, ccc int

		if iscc {
			rcc, scc, pcc, ccc = getContainerCount(ctx, img.ID, apiClient)
		}

		tags := img.RepoTags
		if len(tags) == 0 {
			tags = []string{"<none>:<none>"}
		}
		for _, tag := range tags {
			itag, _ := strings.CutPrefix(img.ID[:19], "sha256:")
			if !iscc {
				fmt.Fprintln(w, itag, "\t", tag, "\t", img.Size/MB)
			} else {
				fmt.Fprintln(w, itag, "\t", tag, "\t", img.Size/MB, "\t", rcc, "\t", scc, "\t", pcc, "\t", ccc)
			}
		}
	}

	w.Flush()

	return nil
}

func getContainerCount(ctx context.Context, id string, apiClient client.APIClient) (int, int, int, int) {
	var scc, rcc, pcc, ccc int
	containers, err := apiClient.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}

	for _, c := range containers {
		if c.ImageID == id {
			if c.State == "running" {
				rcc++
			} else if c.State == "exited" {
				scc++
			} else if c.State == "paused" {
				pcc++
			} else if c.State == "created" {
				ccc++
			}
		}
	}

	return rcc, scc, pcc, ccc
}
