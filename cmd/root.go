package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/starkers/ggg/pkg/git"
	"github.com/starkers/ggg/pkg/logger"
)

var (
	defaultCfg = "~/.config/ggg.toml"
	allArgs    = os.Args
	basePath   string
	cfgFile    string
	timeout    string
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", defaultCfg, "config file")
	rootCmd.PersistentFlags().StringVarP(&basePath, "path", "p", "~/src", "where you would like to keep your git code")
	rootCmd.PersistentFlags().StringVarP(&timeout, "timeout", "t", "1h", "wait this long before assuming git has timed out")
	err := viper.BindPFlag("path", rootCmd.PersistentFlags().Lookup("path"))

	if err != nil {
		logger.Bad(err)
		os.Exit(1)
	}
	viper.SetDefault("path", basePath)
	viper.SetDefault("timeout", timeout)
}

func getExpandedFilePath(input string) (string, error) {
	return homedir.Expand(input)
}

func initConfig() {
	viper.SetConfigType("toml")

	expandedPathToCfg, _ := getExpandedFilePath(defaultCfg)

	viper.SetConfigFile(expandedPathToCfg)
	viper.AddConfigPath(expandedPathToCfg)

	// assuming on config file, its "OK" if we cannot read it (we'll use the defaults/args)
	if cfgFile == defaultCfg {
		err := viper.ReadInConfig()
		if err != nil {
			msg := fmt.Sprintf("first launch.. writing initial config: %s", expandedPathToCfg)
			logger.Info(msg)
			err = viper.SafeWriteConfigAs(expandedPathToCfg)
			if err != nil {
				logger.Bad(err)
				os.Exit(1)
			}
		}
	} else {
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("cannot read config:", err)
			os.Exit(1)
		}
	}

	if didArgLookLikeGit(allArgs) {
		result, err := git.FigureUnixDiskPath(basePath, allArgs[0])
		if err != nil {
			logger.Bad(fmt.Sprintln("doesn't look like a git repo:", allArgs[0]))
		}
		if result == "unknown" {
			logger.Warn("couldn't understand the argument: " + allArgs[0])
		}
	}
}

func didArgLookLikeGit(allArgs []string) bool {
	matchable := []string{
		"git@",
		"https://",
		"http://",
	}
	if len(allArgs) > 1 {
		firstArg := allArgs[1]
		for _, candidateString := range matchable {
			if strings.HasPrefix(firstArg, candidateString) {
				return true
			}
		}
	}
	return false
}
