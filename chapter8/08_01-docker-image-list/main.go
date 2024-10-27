package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

const (
	MB = 1 * 1024 * 1024
)

func main() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug|tabwriter.TabIndent)

	apiClient, err := client.NewClientWithOpts(
		client.FromEnv,
	)
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	images, err := apiClient.ImageList(context.Background(), image.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, "ID", "\t", "Tag", "\t", "Size(MB)")
	for _, img := range images {
		tags := img.RepoTags
		if len(tags) == 0 {
			tags = []string{"<none>:<none>"}
		}
		for _, tag := range tags {
			itag, _ := strings.CutPrefix(img.ID[:19], "sha256:")
			fmt.Fprintln(w, itag, "\t", tag, "\t", img.Size/MB)
		}
	}

	w.Flush()
}
