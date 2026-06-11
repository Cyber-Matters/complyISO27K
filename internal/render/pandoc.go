package render

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
	"github.com/strongdm/comply/internal/config"
)

var pandocArgs = []string{"-f", "markdown+smart", "--toc", "-N", "--template", "templates/default.latex", "-o"}

// pandoc renders outputFilename via either local pandoc or Docker, returning any error.
func pandoc(outputFilename string) error {
	if config.WhichPandoc() == config.UsePandoc {
		return pandocPandoc(outputFilename)
	}
	return dockerPandoc(outputFilename)
}

func dockerPandoc(outputFilename string) (retErr error) {
	pandocCmd := append(pandocArgs, fmt.Sprintf("/source/output/%s", outputFilename), fmt.Sprintf("/source/output/%s.md", outputFilename))
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		return errors.Wrap(err, "unable to read Docker environment")
	}

	pwd, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "unable to get working directory")
	}

	hc := &container.HostConfig{
		Binds: []string{pwd + ":/source"},
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "strongdm/pandoc:edge",
		Cmd:   pandocCmd},
		hc, nil, nil, "")

	if err != nil {
		return errors.Wrap(err, "unable to create Docker container")
	}

	defer func() {
		timeout := 2 * time.Second
		cli.ContainerStop(ctx, resp.ID, &timeout)
		if err := cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{Force: true}); err != nil && retErr == nil {
			retErr = errors.Wrap(err, "unable to remove container")
		}
	}()

	if err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return errors.Wrap(err, "unable to start Docker container")
	}

	chanResult, chanErr := cli.ContainerWait(ctx, resp.ID, "not-running")
	resultValue := <-chanResult

	if resultValue.StatusCode != 0 {
		err = <-chanErr
		return errors.Wrap(err, "error awaiting Docker container")
	}

	_, err = cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return errors.Wrap(err, "error reading Docker container logs")
	}

	if _, err = os.Stat(fmt.Sprintf("output/%s", outputFilename)); err != nil && os.IsNotExist(err) {
		return errors.Wrap(err, "output not generated; verify your Docker image is up to date")
	}

	return nil
}

// 🐼
func pandocPandoc(outputFilename string) error {
	cmd := exec.Command("pandoc", append(pandocArgs, fmt.Sprintf("output/%s", outputFilename), fmt.Sprintf("output/%s.md", outputFilename))...)
	outputRaw, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(outputRaw))
		return errors.Wrap(err, "error calling pandoc")
	}
	return nil
}
