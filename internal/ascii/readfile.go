package ascii

import (
	"bufio"
	"os"
	"strings"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-fs/pkg/model"
)

func (a *Ascii) ReadFont(char rune) ([]string, *m.Error) {
	if char < ' ' || char > '~' {
		return nil, &m.Error{
			Error:  "Invalid character.",
			Detail: "Found an invalid character.\nUnable to transform invalid character",
		}
	}

	path := ""

	switch strings.ToLower(a.Banner) {
	case "shadow":
		path = "fonts/shadow.txt"
	case "standard":
		path = "fonts/standard.txt"
	case "tinkertoy":
		path = "fonts/tinkertoy.txt"
	default:
		return nil, &m.Error{
			Error:  "Invalid font.",
			Detail: "The font you gave is not supported by our system.\nSupported fonts => Standard, Shadow & Tinkertoy",
		}
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, &m.Error{
			Error:  "File read error.",
			Detail: err.Error(),
		}
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	startLine := ((char - ' ') * 9) + 1
	endLine := startLine + 8
	currentLine := 0

	var output []string

	for scanner.Scan() {
		if currentLine >= int(startLine) && currentLine <= int(endLine) {
			output = append(output, scanner.Text())
		} else if currentLine > int(endLine) {
			break
		}
		currentLine += 1
	}

	if err := scanner.Err(); err != nil {
		return nil, &m.Error{
			Error:  "File read error.",
			Detail: err.Error(),
		}
	}
	return output, nil
}
