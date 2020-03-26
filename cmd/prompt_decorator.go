package cmd

import (
	"fmt"

	"github.com/containerd/console"
	"github.com/fatih/color"
	"github.com/nlowe/kubectl-prompt_decorator/decorator"
	"github.com/nlowe/kubectl-prompt_decorator/decorator/theme"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

type options struct {
	disableColors bool
	theme         string

	persistentFlags *genericclioptions.ConfigFlags
}

func NewCmdPromptDecorator() *cobra.Command {
	opts := &options{
		persistentFlags: genericclioptions.NewConfigFlags(true),
	}

	cmd := &cobra.Command{
		Use:   "prompt-decorator [flags]",
		Short: "Decorate your shell prompt with your current context information",
		Example: `$ kubectl prompt-decorator
[☸ my-cluster:my-default-namespace]`,
		SilenceUsage: true,
		RunE: func(_ *cobra.Command, _ []string) (err error) {
			cfg, err := genericclioptions.NewConfigFlags(true).ToRawKubeConfigLoader().RawConfig()
			if err != nil {
				return err
			}

			if opts.disableColors {
				opts.theme = theme.None
			}

			con := console.Current()
			defer func() {
				_ = con.Reset()
			}()

			if err := con.SetRaw(); err != nil {
				return err
			}

			fmt.Print(decorator.Format(cfg, opts.theme))
			return nil
		},
	}

	flags := cmd.Flags()
	flags.BoolVar(&color.NoColor, "disable-colors", false, "Disable color ANSI Escape Sequences")
	flags.StringVar(&opts.theme, "theme", "default", "The theme to use")

	return cmd
}
