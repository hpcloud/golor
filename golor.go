// 256 color terminal printing in Go.
package golor

import (
	"crypto/sha1"
	"fmt"
)

// 16-colors
var RED, GREEN, BLUE, CYAN, MAGENTA, YELLOW, BLACK, WHITE int
var R, G, C, M, Y, W int

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

// RGB returns the color code corresponding a RGB value set. Arguments
// must range from 0 to 5.
func RGB(red, green, blue int) int {
	return 16 + (red * 36) + (green * 6) + blue
}

// Assign and return an unique color for this string. Automatically
// avoid black/white colors.
func AssignColor(s string) int {
	maxColor := 211 // prevent whitish colors; use prime number for mod.
	minColor := 17  // avoid 16 colors, including the black.
	fg := minColor + stringId(s, maxColor-minColor+1)
	return fg
}

func stringId(s string, mod int) int {
	h := sha1.New()
	h.Write([]byte(s))
	sum := 0
	for _, n := range h.Sum(nil) {
		sum += int(n)
	}
	return sum % mod
}

func init() {
	RED = RGB(5, 0, 0)
	GREEN = RGB(0, 5, 0)
	BLUE = RGB(0, 0, 5)
	CYAN = RGB(0, 5, 5)
	MAGENTA = RGB(5, 0, 5)
	YELLOW = RGB(5, 5, 0)
	BLACK = 0
	WHITE = RGB(5, 5, 5)

	R = RED
	G = GREEN
	C = CYAN
	M = MAGENTA
	Y = YELLOW
	W = WHITE
}
