package VIC2

import (
	"github.com/gdamore/tcell"
)

//TODO use github.com/fatih/color

func getColor(value byte) (st tcell.Color) {
	value = value & 15
	switch value {
	case 0:
		return tcell.ColorBlack
	case 1:
		return tcell.ColorWhite
	case 2:
		return tcell.ColorDarkRed
	case 3:
		return tcell.ColorLightCyan
	case 4:
		return tcell.ColorMaroon
	case 5:
		return tcell.ColorDarkGreen
	case 6:
		return tcell.ColorDarkBlue
	case 7:
		return tcell.ColorYellow
	case 8:
		return tcell.ColorYellowGreen
	case 9:
		return tcell.ColorDarkMagenta
	case 10:
		return tcell.ColorRed
	case 11:
		return tcell.ColorDarkGray
	case 12:
		return tcell.ColorGray
	case 13:
		return tcell.ColorGreen
	case 14:
		return tcell.ColorBlue
	case 15:
		return tcell.ColorDarkCyan
	}

	return tcell.ColorWhite
}
