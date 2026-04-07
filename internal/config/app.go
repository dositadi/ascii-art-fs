package config

import (
	"fmt"
	"os"

	s "acad.learn2earn.ng/git/dositadi/ascii-art-fs/internal/ascii"
	m "acad.learn2earn.ng/git/dositadi/ascii-art-fs/pkg/model"
)

type App struct{}

type Sketch interface {
	Splitter() []string
	ReadFont(char rune) ([]string, *m.Error)
	Transform(input []string) [][][]string
}

func (a *App) GetInput() (input, banner string) {
	if len(os.Args) == 1 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
		return
	}

	args := os.Args[1:]

	switch len(args) {
	case 0:
		fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
		return
	case 1:
		input = args[0]
		banner = "standard"
	case 2:
		input = args[0]
		banner = args[1]
	default:
		fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
		return
	}
	return input, banner
}

func (a *App) Draw() {
	input, banner := a.GetInput()
	sketch := s.Ascii{Input: input, Banner: banner}

	trimmed := sketch.Splitter()
	fmt.Println(trimmed, len(trimmed))
	_, err := sketch.ReadFont('a')
	if err != nil {
		panic(err)
	}
}

func (a *App) Run() {
	a.Draw()
}
