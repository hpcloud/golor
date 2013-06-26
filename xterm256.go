package golor

import (
	"fmt"
)

type Xterm256 struct {
}

func (t *Xterm256) Colorize(s string, fg int, bg int) string {
	if fg > 255 || bg > 255 {
		panic("invalid color index")
	}
	color := ""
	if fg > -1 {
		color += fmt.Sprintf("\033[38;5;%dm", fg)
	}
	if bg > -1 {
		color += fmt.Sprintf("\033[48;5;%dm", bg)
	}
	reset := "\033[0m"
	return color + s + reset
}

// RGB returns the color code corresponding a RGB value set. Arguments
// must range from 0 to 5. Returned values range from 16 (black) to
// 231 (white).
func (t *Xterm256) RGB(red, green, blue int) int {
	return 16 + (red * 36) + (green * 6) + blue
}

func (t *Xterm256) AssignColor(s string) int {
	// Note: MAX-2 == 229 (in 256-colors terminal), which is also a
	// prime number that is suitable for the modulo in stringId()
	maxColor := MAX - 2 // prevent whitish colors
	minColor := MIN + 4 // prevent black and dark blue.

	fg := minColor + stringId(s, maxColor-minColor+1)
	return fg
}
