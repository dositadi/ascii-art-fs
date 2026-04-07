package ascii

import (
	"bufio"
	"fmt"
	"os"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-fs/pkg/model"
)

func (a *Ascii) ReadFont(char rune, font string) ([]string, *m.Error) {
	fmt.Println("Entered")
	path := ""

	switch font {
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

	startLine := (char - ' ') * 9
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
