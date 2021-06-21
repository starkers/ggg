package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/starkers/ggg/pkg/logger"
	"github.com/starkers/ggg/pkg/shell"

	"github.com/spf13/cobra"
	"github.com/starkers/ggg/pkg/git"
)

var (
	repoURLFlag string
)

func init() {
	rootCmd.AddCommand(workCmd)
	workCmd.Flags().StringVarP(&repoURLFlag, "repo", "r", "", "remote repo url, eg git@github.com:foo/bar.git")
	err := workCmd.MarkFlagRequired("repo")
	if err != nil {
		return
	}
}

var workCmd = &cobra.Command{
	Use:   "work",
	Short: "do work",
	Long:  `do work n stuff`,
	Run: func(cmd *cobra.Command, args []string) {
		doWork()
	},
}

func doWork() {
	//logger.Info("doing work for: " + repoURLFlag)
	result, err := git.FigureUnixDiskPath(basePath, repoURLFlag)
	if err != nil {
		logger.Bad(err)
		os.Exit(1)
	}
	if directoryExists(result) {
		logger.Good("cd " + result)
		return
	}

	// try to make the dir above (so we can do a clone afterwards)

	splitBySlash := strings.Split(result, "/")
	slashCount := len(splitBySlash)
	subDir := strings.Join(splitBySlash[0:slashCount-1], "/")
	if directoryExists(subDir) {
		cmd := fmt.Sprintf("cd %s && git clone %s", subDir, repoURLFlag)
		// logger.Good(fmt.Sprintf("Running 'git clone %s' -> %s", repoURLFlag, result))
		err := shell.Run(cmd, 30*time.Second)
		if err != nil {
			logger.Bad(err)
		}
	} else {
		cmd := fmt.Sprintf("mkdir -p %s", subDir)
		logger.Good("making dir: " + subDir)
		err := shell.Run(cmd, 2*time.Second)
		if err != nil {
			logger.Bad(err)
		}
		if directoryExists(subDir) {
			cmd := fmt.Sprintf("cd %s && git clone %s", subDir, repoURLFlag)
			// logger.Good(fmt.Sprintf("Running 'git clone %s' -> %s", repoURLFlag, result))
			err := shell.Run(cmd, 30*time.Second)
			if err != nil {
				logger.Bad(err)
			}
		}
	}
}

func directoryExists(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		logger.Warn(err)
		return false
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logger.Warn(err)
			//log.Debug(err)
		}
	}(file)

	fStat, err := file.Stat()
	if err != nil {
		logger.Warn(err)
		//log.Debug(err)
		return false
	}
	return fStat.IsDir()
}
