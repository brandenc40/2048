package main

import (
	"flag"

	"github.com/brandenc40/2048/game"
	"github.com/brandenc40/2048/terminalui"
)

func main() {
	var output string
	flag.StringVar(&output, "output", "rgb",
		`Output mode use for displaying colors in the terminal. Options are "rgb", "256", and "normal". 
If you are not able to see the game board, your terminal most likely does not support "rgb". 
In that case please use "256", or "normal".`)

	flag.Parse()

	terminalui.Run(
		game.NewController(),
		parseOutModeOption(output),
	)
}

func parseOutModeOption(output string) terminalui.Option {
	switch output {
	case "normal":
		return terminalui.WithOutputMode(terminalui.OutputModeNormal)
	case "256":
		return terminalui.WithOutputMode(terminalui.OutputMode256)
	case "rgb":
		return terminalui.WithOutputMode(terminalui.OutputModeRGB)
	default:
		return terminalui.WithOutputMode(terminalui.OutputModeRGB)
	}
}
