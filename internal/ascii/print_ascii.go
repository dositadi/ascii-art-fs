package ascii

import (
	"fmt"
	"slices"
	"strings"
	"time"
)

func (a *Ascii) PrintAscii(input [][][]string) {
	for i := 0; i < len(input); i++ {
		currentWord := input[i]

		if slices.Compare(currentWord[0], []string{""}) == 0 {
			fmt.Println()
			continue
		}

		for j := 0; j < 8; j++ {
			var combinedString strings.Builder
			for k := range currentWord {
				if k != len(currentWord)-1 {
					combinedString.WriteString(currentWord[k][j] + " ")
				} else {
					combinedString.WriteString(currentWord[k][j])
				}
			}
			fmt.Println(combinedString.String())
			time.Sleep(1 * time.Second)
		}
	}
}
