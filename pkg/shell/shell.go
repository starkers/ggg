package shell

import (
	"bytes"
	"context"
	"github.com/starkers/ggg/pkg/logger"
	"os"
	"os/exec"
	"time"
)

func Run(
	command string,
	timeout time.Duration,
) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var outBuffer bytes.Buffer
	logger.Good("running: " + command)

	runMe := []string{"-c", command}
	cmd := exec.CommandContext(ctx, "sh", runMe...)

	cmd.Stdout = &outBuffer
	cmd.Stderr = &outBuffer

	start := time.Now()

	err := cmd.Run()
	duration := time.Since(start)

	logger.Good("command took: " + duration.String())
	if err != nil {
		logger.Bad(err)
		logger.Bad(cmd.Stdout)
		os.Exit(5)
	}
	return nil

	// TODO: turn on timeout+context checking

	//// Check if the ctx timeout was hit first
	//if ctx.Err() == context.DeadlineExceeded {
	//	log.Infow("command timed out.. not ",
	//		"error", err,
	//		"contextErr", ctx.Err(),
	//		"output", string(outBuffer.Bytes()),
	//		"duration", duration,
	//	)
	//	return ctx.Err()
	//}
	//
	//if err != nil {
	//	log.Errorw("Non-zero exit code",
	//		"output", string(outBuffer.Bytes()),
	//		"error", err,
	//	)
	//}
	//return err

}
