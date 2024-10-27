package main

import (
	"context"
	"encoding/json"
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

type Metadata struct {
	SchemaVersion    string `json:"SchemaVersion"`
	Vendor           string `json:"Vendor"`
	Version          string `json:"Version"`
	ShortDescription string `json:"ShortDescription"`
	URL              string `json:"URL"`
}

var metadata = Metadata{
	SchemaVersion:    "0.1.0",
	Version:          "0.1.0",
	Vendor:           "Cloud Native Central",
	ShortDescription: "Provides a list of container images on the host",
	URL:              "github.com/rchaganti/docker-imagelist",
}

func main() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug|tabwriter.TabIndent)

	if len(os.Args) > 1 && os.Args[1] == "docker-cli-plugin-metadata" {
		b, err := json.Marshal(metadata)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(b))
	} else {
		apiClient, err := client.NewClientWithOpts(
			client.WithAPIVersionNegotiation(),
			client.FromEnv,
		)
		if err != nil {
			panic(err)
		}
		defer apiClient.Close()

		images, err := apiClient.ImageList(context.Background(), image.ListOptions{All: true})
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
}
