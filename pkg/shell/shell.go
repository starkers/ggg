package shell

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/starkers/ggg/pkg/logger"
)

func Run(
	command string,
) error {

	// adding a huge 15 minute timeout
	// TODO: ensure timeout is handled correctly and provide configuration options
	timeout := 15 * time.Minute
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	logger.Good("running: sh " + command)

	// assemble with a shell, eg: "sh -c runMe..."
	runMe := []string{"-c", command}
	cmd := exec.CommandContext(ctx, "sh", runMe...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	// redirect stdout -> stderr
	cmd.Stderr = cmd.Stdout

	// start running it
	if err = cmd.Start(); err != nil {
		return err
	}
	// Get real-time output, break when stdoutRead has its first err (typically some kind of EOF)
	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)
		fmt.Print(string(tmp))
		if err != nil {
			break
		}
	}

	// block and wait for the cmd to complete
	if err = cmd.Wait(); err != nil {
		return err
	}

	return nil

}
