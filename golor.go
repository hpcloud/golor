// 256 color terminal printing in Go.
package golor

import (
	"fmt"
)

func Colorize(s string, fg int, bg int) string {
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

// RGB returns the color code corresponding a RGB value set.
func RGB(red, green, blue int) int {
	return 16 + (red * 36) + (green * 6) + blue
}
