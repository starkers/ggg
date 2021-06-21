package cmd

import (
	"github.com/spf13/cobra"
	"github.com/starkers/ggg/pkg/git"
	"github.com/starkers/ggg/pkg/logger"
	"os"
)

func init() {
	rootCmd.AddCommand(destCmd)
	destCmd.Flags().StringVarP(&repoURLFlag, "repo", "r", "", "remote repo url, eg git@github.com:foo/bar.git")
	err := destCmd.MarkFlagRequired("repo")
	if err != nil {
		return
	}
}

var destCmd = &cobra.Command{
	Use:   "dest",
	Short: "do dest",
	Long:  `do dest n stuff`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := git.FigureUnixDiskPath(basePath, repoURLFlag)
		if err != nil {
			logger.Bad(err)
			os.Exit(1)
		}
		logger.Raw(result)
	},
}
