package main

import (
	"os"

	"github.com/mattn/go-colorable"
	"github.com/nlowe/kubectl-prompt_decorator/cmd"
	"github.com/spf13/pflag"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func main() {
	pflag.CommandLine = pflag.NewFlagSet("kubectl-prompt_decorator", pflag.ExitOnError)

	streams := genericclioptions.IOStreams{
		In:     os.Stdin,
		Out:    colorable.NewColorableStdout(),
		ErrOut: colorable.NewColorableStderr(),
	}

	if err := cmd.NewCmdPromptDecorator(streams).Execute(); err != nil {
		os.Exit(1)
	}
}
