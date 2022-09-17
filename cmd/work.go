package cmd

import (
	"fmt"
	"github.com/starkers/ggg/pkg/logger"
	"github.com/starkers/ggg/pkg/shell"
	"os"
	"strings"

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
	Long:  `do work and stuff`,
	Run: func(cmd *cobra.Command, args []string) {
		doWork()
	},
}

func doWork() {
	// logger.Info("doing work for: " + repoURLFlag)
	result, err := git.FigureUnixDiskPath(basePath, repoURLFlag)
	if err != nil {
		logger.Bad(err)
		os.Exit(1)
	}
	if directoryExists(result) {
		logger.Good("cd " + result)
		return
	}

	splitBySlash := strings.Split(result, "/")
	slashCount := len(splitBySlash)
	subDir := strings.Join(splitBySlash[0:slashCount-1], "/")
	cloneCmd := "git clone --progress " + repoURLFlag
	if directoryExists(subDir) {
		cmd := fmt.Sprintf("cd %s && %s", subDir, cloneCmd)
		err := shell.Run(cmd)
		if err != nil {
			logger.Bad(err)
		}
	} else {
		cmd := fmt.Sprintf("mkdir -pv %s", subDir)
		logger.Good("making dir: " + subDir)
		err := shell.Run(cmd)
		if err != nil {
			logger.Bad(err)
		}
		if directoryExists(subDir) {
			cmd := fmt.Sprintf("cd %s && %s", subDir, cloneCmd)
			err := shell.Run(cmd)
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
		}
	}(file)

	fStat, err := file.Stat()
	if err != nil {
		logger.Warn(err)
		return false
	}
	return fStat.IsDir()
}
