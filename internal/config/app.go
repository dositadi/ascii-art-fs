package config

import (
	"fmt"
	"os"

	s "acad.learn2earn.ng/git/dositadi/ascii-art-fs/internal/ascii"
	m "acad.learn2earn.ng/git/dositadi/ascii-art-fs/pkg/model"
)

type App struct{}

const (
	Usage = "Usage: go run . [STRING] [BANNER]\nEX: go run . something standard"
)

type Sketch interface {
	Splitter() []string
	ReadFont(char rune) ([]string, *m.Error)
	Transform(input []string) [][][]string
	PrintAscii(input [][][]string)
}

func (a *App) GetInput() (input, banner string) {
	if len(os.Args) == 1 {
		PrintError(&m.Error{Error: "Invalid command", Detail: "You did not enter the text and your choice banner"})
		return
	}

	args := os.Args[1:]

	switch len(args) {
	case 0:
		PrintError(&m.Error{Error: "Invalid command", Detail: "You did not enter the text and your choice banner"})
		return
	case 1:
		input = args[0]
		banner = "standard"
	case 2:
		input = args[0]
		banner = args[1]
	default:
		PrintError(&m.Error{Error: "Superflous input", Detail: "Only two inputs are required (the <text> and <banner>)"})
		return
	}
	return input, banner
}

func (a *App) Draw() {
	input, banner := a.GetInput()

	if input == "" || banner == "" {
		return
	}

	sketch := s.Ascii{Input: input, Banner: banner}

	trimmed := sketch.Splitter()

	formattedAscii, err2 := sketch.Transform(trimmed)
	if err2 != nil {
		PrintError(err2)
		return
	}

	sketch.PrintAscii(formattedAscii)
}

func (a *App) Run() {
	a.Draw()
}

func PrintError(err *m.Error) {
	message := fmt.Sprintf("Error: %s\nDetail: %s", err.Error, err.Detail)

	fmt.Println(Usage, "\n\n", message)
}
