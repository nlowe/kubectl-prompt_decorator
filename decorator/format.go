package decorator

import (
	"github.com/nlowe/kubectl-prompt_decorator/decorator/theme"
	"k8s.io/client-go/tools/clientcmd/api"
)

func Format(cfg api.Config, themeName string) string {
	if cfg.CurrentContext == "" {
		return ""
	}

	ctx, found := cfg.Contexts[cfg.CurrentContext]
	if !found {
		return ""
	}

	ns := "default"
	if found && ctx.Namespace != "" {
		ns = ctx.Namespace
	}

	t := theme.Lookup(themeName)
	return t.Brace("[") + t.Icon("☸️") + " " + t.Cluster("%s", ctx.Cluster) + t.Separator(":") + t.Namespace("%s", ns) + t.Brace("]")
}
