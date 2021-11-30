package terminalui

import "github.com/nsf/termbox-go"

type Option interface {
	apply(ui *ui)
}

func WithOutputMode(mode OutputMode) Option {
	return outputModeOption{mode: mode}
}

type outputModeOption struct {
	mode OutputMode
}

func (o outputModeOption) apply(ui *ui) {
	switch o.mode {
	case OutputMode256:
		termbox.SetOutputMode(termbox.Output256)
		ui.colorPalate = mode256Palate()
	case OutputModeRGB:
		termbox.SetOutputMode(termbox.OutputRGB)
		ui.colorPalate = modeRGBPalate()
	case OutputModeNormal:
		termbox.SetOutputMode(termbox.OutputNormal)
		ui.colorPalate = normalPalate()
	default:
		panic("WithOutputMode: invalid output mode")
	}
}
