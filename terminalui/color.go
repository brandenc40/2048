package terminalui

import "github.com/nsf/termbox-go"

type OutputMode int8

const (
	OutputModeNormal OutputMode = iota
	OutputMode256
	OutputModeRGB
)

type colorPalate struct {
	values      map[uint16]termbox.Attribute
	valueText   termbox.Attribute
	empty       termbox.Attribute
	border      termbox.Attribute
	score       termbox.Attribute
	guide       termbox.Attribute
	overlayText termbox.Attribute
	overlayBg   termbox.Attribute
}

func normalPalate() colorPalate {
	return colorPalate{
		values: map[uint16]termbox.Attribute{
			2:    termbox.ColorLightGray,
			4:    termbox.ColorLightRed,
			8:    termbox.ColorRed,
			16:   termbox.ColorLightGreen,
			32:   termbox.ColorGreen,
			64:   termbox.ColorLightBlue,
			128:  termbox.ColorBlue,
			256:  termbox.ColorLightYellow,
			512:  termbox.ColorYellow,
			1024: termbox.ColorCyan,
			2048: termbox.ColorMagenta,
		},
		empty:       termbox.ColorWhite,
		border:      termbox.ColorDarkGray,
		score:       termbox.ColorGreen | termbox.AttrBold,
		guide:       termbox.ColorWhite | termbox.AttrBold,
		overlayText: termbox.ColorWhite | termbox.AttrBold,
		overlayBg:   termbox.ColorBlack,
		valueText:   termbox.ColorBlack,
	}
}

func mode256Palate() colorPalate {
	const colorMod = 1 << 32

	return colorPalate{
		values: map[uint16]termbox.Attribute{
			2:    termbox.ColorLightGray,
			4:    10 + colorMod,
			8:    11 + colorMod,
			16:   12 + colorMod,
			32:   13 + colorMod,
			64:   14 + colorMod,
			128:  15 + colorMod,
			256:  16 + colorMod,
			512:  6 + colorMod,
			1024: 5 + colorMod,
			2048: 4 + colorMod,
		},
		empty:       termbox.ColorWhite,
		border:      termbox.ColorDarkGray,
		score:       termbox.ColorGreen | termbox.AttrBold,
		guide:       termbox.ColorWhite | termbox.AttrBold,
		overlayText: termbox.ColorWhite | termbox.AttrBold,
		overlayBg:   termbox.ColorBlack,
		valueText:   termbox.ColorBlack,
	}
}

func modeRGBPalate() colorPalate {
	return colorPalate{
		values: map[uint16]termbox.Attribute{
			2:    termbox.RGBToAttribute(235, 228, 219),
			4:    termbox.RGBToAttribute(234, 225, 204),
			8:    termbox.RGBToAttribute(233, 180, 129),
			16:   termbox.RGBToAttribute(232, 154, 108),
			32:   termbox.RGBToAttribute(231, 132, 103),
			64:   termbox.RGBToAttribute(229, 105, 72),
			128:  termbox.RGBToAttribute(232, 209, 128),
			256:  termbox.RGBToAttribute(232, 205, 114),
			512:  termbox.RGBToAttribute(231, 202, 101),
			1024: termbox.RGBToAttribute(230, 197, 90),
			2048: termbox.RGBToAttribute(230, 196, 79),
		},
		valueText:   termbox.RGBToAttribute(32, 32, 31) | termbox.AttrBold,
		empty:       termbox.RGBToAttribute(202, 193, 181),
		border:      termbox.RGBToAttribute(185, 173, 161),
		score:       termbox.RGBToAttribute(71, 155, 95) | termbox.AttrBold,
		guide:       termbox.RGBToAttribute(250, 248, 240) | termbox.AttrBold,
		overlayText: termbox.RGBToAttribute(250, 248, 240) | termbox.AttrBold,
		overlayBg:   termbox.RGBToAttribute(32, 32, 31),
	}
}
