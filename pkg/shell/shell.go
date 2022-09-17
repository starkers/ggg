package shell

import (
	"context"
	"fmt"
	str2duration "github.com/xhit/go-str2duration/v2"
	"os"
	"os/exec"

	"github.com/spf13/viper"
	"github.com/starkers/ggg/pkg/logger"
)

func Run(command string) error {
	timeoutFromConfig := viper.Get("timeout")
	msg := fmt.Sprintf("using git timeout: %s", timeoutFromConfig)
	logger.Info(msg)
	durationFromString, err := str2duration.ParseDuration(fmt.Sprintf("%s", timeoutFromConfig))
	if err != nil {
		msg := fmt.Sprintf("couldn't understand the timeout '%s'.. please use something like '1h', or '2d1h10m'", timeoutFromConfig)
		logger.Bad(msg)
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), durationFromString)
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
