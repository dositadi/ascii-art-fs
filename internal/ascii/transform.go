package ascii

import (
	m "acad.learn2earn.ng/git/dositadi/ascii-art-fs/pkg/model"
)

func (a *Ascii) Transform(input []string) ([][][]string, *m.Error) {
	var output [][][]string

	for _, word := range input {
		currentWord := word
		var wordAscii [][]string

		if currentWord == "" {
			wordAscii = append(wordAscii, []string{"\n"})
			output = append(output, wordAscii)
			continue
		}

		for _, char := range currentWord {
			charAscii, err := a.ReadFont(char)
			if err != nil {
				return nil, err
			}

			wordAscii = append(wordAscii, charAscii)
			output = append(output, wordAscii)
		}
	}
	return output, nil
}
