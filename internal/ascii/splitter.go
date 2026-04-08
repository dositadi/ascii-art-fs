package ascii

import (
	"strings"
)

func (a *Ascii) Splitter() []string {
	output := strings.Split(a.Input, "\\n")
	return output
}
