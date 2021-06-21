package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ggg",
	Args:  cobra.ArbitraryArgs,
	Short: "ggg is a simple (recursive) git cloneing tool",
	Long:  `GGG will get you to your git repos fast`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
