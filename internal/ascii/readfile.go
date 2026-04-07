package ascii

import (
	"bufio"
	"os"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-fs/pkg/model"
)

func (a *Ascii) ReadFont(char rune) ([]string, *m.Error) {
	path := ""

	switch a.Banner {
	case "shadow":
		path = "fonts/shadow.txt"
	case "standard":
		path = "fonts/standard.txt"
	case "tinkertoy":
		path = "fonts/tinkertoy.txt"
	default:
		return nil, &m.Error{
			Error:  "Invalid font.",
			Detail: "The font you gave is not supported by our system.",
		}
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, &m.Error{
			Error:  "File read error.",
			Detail: err.Error(),
		}
	}

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
