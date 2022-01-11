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
	cfgFile    string
	basePath   string
	allArgs    = os.Args
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", defaultCfg, "config file")
	rootCmd.PersistentFlags().StringVarP(&basePath, "path", "p", "~/src", "where you would like to keep your git code")
	err := viper.BindPFlag("path", rootCmd.PersistentFlags().Lookup("path"))

	if err != nil {
		logger.Bad(err)
		os.Exit(1)
	}
	viper.SetDefault("path", basePath)
}

func getExpandedFilePath(input string) (string, error) {
	return homedir.Expand(input)
}

func initConfig() {
	viper.SetConfigType("toml")

	expandedPathToCfg, _ := getExpandedFilePath(defaultCfg)

	viper.SetConfigFile(expandedPathToCfg)
	viper.AddConfigPath(expandedPathToCfg)

	/* on first run write out a default config file if one does not exist already */
	if cfgFile == defaultCfg {
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Printf("# no config found, writing initial config to: %s\n", expandedPathToCfg)
			err = viper.SafeWriteConfigAs(expandedPathToCfg)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

	} else {

		// if the user has a config file, we must be able to read it

		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("Can't read config:", err)
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
