package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/starkers/ggg/pkg/logger"
	"os"
)

func init() {
	rootCmd.AddCommand(hookCmd)
	hookCmd.AddCommand(forBash)
	hookCmd.AddCommand(forZSH)
	hookCmd.AddCommand(forFish)
}

var hookCmd = &cobra.Command{
	Use:   "hook",
	Short: "Export shell bindings",
	Long:  `Run this in your shells init.. eg: ~/.zshrc or ~/.bashrc`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// basicExporter()
	},
}

var forFish = &cobra.Command{
	Use:   "fish",
	Short: "fish",
	Long:  `for fish`,
	Run: func(cmd *cobra.Command, args []string) {
		hookFish()
	},
}
var forBash = &cobra.Command{
	Use:   "bash",
	Short: "bash",
	Long:  `for bash`,
	Run: func(cmd *cobra.Command, args []string) {
		hookGeneric("bash")
	},
}

var forZSH = &cobra.Command{
	Use:   "zsh",
	Short: "zsh",
	Long:  `for zsh`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("echo for zsh")
		hookGeneric("zsh")
	},
}

// same looking config hook for zsh or bash
func hookGeneric(shell string) {
	self := getSelf()
	data := fmt.Sprintf(`

# vi: ft=%s
# config for fish by 'ggg'

function ggg(){
  local GGG_BIN="%s"
  if [ -f "${GGG_BIN}" ]; then
    if "${GGG_BIN}" work -r "${1}" ; then
      local DEST="$("${GGG_BIN}" dest -r "${1}")"
      cd "${DEST}"
    fi
  fi
}

`, shell, self)

	fmt.Println(data)
}

// fish only hook
func hookFish() {
	self := getSelf()

	data := fmt.Sprintf(`
# vs: ft=fish
# config for fish by 'ggg'

set GGG_BIN "%s"
if test -f $GGG_BIN
  function ggg
    "$GGG_BIN" work -r "$argv[1]"
    cd ( "$GGG_BIN" dest -r "$argv[1]" )
  end
end
`, self)

	logger.Raw(data)
}

func getSelf() string {
	self, err := os.Executable()
	if err != nil {
		logger.Bad(err)
		os.Exit(1)
	}
	return self
}
