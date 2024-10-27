package helper

import (
	"context"
	"encoding/json"
	"io"
	"strings"

	"github.com/docker/cli/cli/command"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type Container struct {
	ID         string
	Name       string
	Image      string
	State      string
	IsSelected bool
}

func (c Container) FilterValue() string {
	return c.State
}

func (c Container) Title() string {
	return c.ID[:12]
}

func (c Container) Description() string {
	return c.State
}

func GetContainers(dockerCli command.Cli) ([]Container, error) {
	var containers []Container

	apiClient := dockerCli.Client()
	defer apiClient.Close()

	ctx := context.Background()

	c, err := apiClient.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return nil, err
	}

	for i, container := range c {
		image, err := getImageName(apiClient, container.ImageID)
		if err != nil {
			return nil, err
		}
		containers = append(containers, Container{
			ID:    container.ID,
			Name:  strings.TrimPrefix(container.Names[0], "/"),
			Image: image,
			State: container.State,
		})
		if i == 0 {
			containers[0].IsSelected = true
		}
	}

	return containers, nil
}

func getImageName(apiClient client.APIClient, imageID string) (string, error) {
	image, _, err := apiClient.ImageInspectWithRaw(context.Background(), imageID)
	if err != nil {
		return "", err
	}

	return image.RepoTags[0], nil
}

func GetContainerLogs(dockerCli command.Cli, containerID string) (string, error) {
	apiClient := dockerCli.Client()
	defer apiClient.Close()

	ctx := context.Background()

	logsReader, err := apiClient.ContainerLogs(
		ctx,
		containerID,
		container.LogsOptions{
			ShowStdout: true,
			ShowStderr: true,
		},
	)
	if err != nil {
		return "", err
	}
	defer logsReader.Close()

	logs, err := io.ReadAll(logsReader)
	if err != nil {
		return "", err
	} else if strings.Trim(string(logs), " ") == "" {
		return "No logs found", nil
	}

	return string(logs), nil
}

func GetContainerInspect(dockerCli command.Cli, containerID string) (string, error) {
	apiClient := dockerCli.Client()
	defer apiClient.Close()

	ctx := context.Background()

	_, rawJson, err := apiClient.ContainerInspectWithRaw(ctx, containerID, true)
	if err != nil {
		return "", err
	}

	var obj map[string]interface{}
	err = json.Unmarshal(rawJson, &obj)
	if err != nil {
		return "", err
	}

	jsonStr, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonStr), nil

}

func StopContainer(dockerCli command.Cli, containerID string) error {
	apiClient := dockerCli.Client()
	defer apiClient.Close()

	ctx := context.Background()

	err := apiClient.ContainerStop(ctx, containerID, container.StopOptions{})
	if err != nil {
		return err
	}
	return nil
}
