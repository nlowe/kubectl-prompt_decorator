package theme

import (
	"fmt"

	"github.com/fatih/color"
)

type Component func(string, ...interface{}) string

type Theme struct {
	Brace     Component
	Icon      Component
	Cluster   Component
	Separator Component
	Namespace Component
}

var themes = map[string]Theme{
	"none": {
		Brace:     fmt.Sprintf,
		Icon:      fmt.Sprintf,
		Cluster:   fmt.Sprintf,
		Separator: fmt.Sprintf,
		Namespace: fmt.Sprintf,
	},
}

func Lookup(name string) Theme {
	t, ok := themes[name]
	if !ok {
		return Theme{
			Brace:     color.New(color.Bold, color.FgYellow).SprintfFunc(),
			Icon:      color.New(color.FgMagenta).SprintfFunc(),
			Cluster:   color.New(color.FgHiCyan).SprintfFunc(),
			Separator: color.New(color.FgWhite).SprintfFunc(),
			Namespace: color.New(color.FgHiCyan).SprintfFunc(),
		}
	}

	return t
}
