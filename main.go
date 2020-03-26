package main

import (
	"os"

	"github.com/nlowe/kubectl-prompt_decorator/cmd"
	"github.com/spf13/pflag"
)

func main() {
	pflag.CommandLine = pflag.NewFlagSet("kubectl-prompt_decorator", pflag.ExitOnError)

	if err := cmd.NewCmdPromptDecorator().Execute(); err != nil {
		os.Exit(1)
	}
}
