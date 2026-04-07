package config

import (
	"fmt"
	"os"
)

type App struct{}

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

func (a *App) Run() {
	input, banner := a.GetInput()
	fmt.Println("input: ", input, " banner: ", banner)
}
