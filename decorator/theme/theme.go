package theme

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

const (
	sequencePreamble = "\\[\033["
	sequenceTrailer  = "m\\]"
	resetSequence    = "\001\033[0;0m\002"

	None = "none"
)

type Component func(string, ...interface{}) string

func NewComponent(opts ...color.Attribute) Component {
	return func(s string, args ...interface{}) string {
		seq := strings.Builder{}
		seq.WriteString(sequencePreamble)
		for i, o := range opts {
			seq.WriteString(strconv.Itoa(int(o)))
			if i < len(opts)-1 {
				seq.WriteRune(';')
			}
		}
		seq.WriteString(sequenceTrailer)

		return fmt.Sprintf("%s%s%s", seq.String(), fmt.Sprintf(s, args...), resetSequence)
	}
}

type Theme struct {
	Brace     Component
	Icon      Component
	Cluster   Component
	Separator Component
	Namespace Component
}

var themes = map[string]Theme{
	None: {
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
			Brace:     NewComponent(color.Bold, color.FgYellow),
			Icon:      NewComponent(color.FgMagenta),
			Cluster:   NewComponent(color.FgHiCyan),
			Separator: NewComponent(color.FgWhite),
			Namespace: NewComponent(color.FgHiCyan),
		}
	}

	return t
}
