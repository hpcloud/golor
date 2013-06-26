package golor

import (
	"fmt"
)

// Legacy terminal supporting only 16 colors.
type Legacy struct {
}

func (t *Legacy) Colorize(s string, fg int, bg int) string {
	color := ""
	attr := 0
	var ok bool

	// fmt.Printf("Orig =  %d, %d\n", fg, bg)

	if fg > 15 || bg > 15 {
		panic("invalid color index")
	}

	if fg, ok = checkBold(fg); ok {
		attr = 1 // bold code
	}
	if bg, ok = checkBold(bg); ok {
		attr = 1 // bold code
	}

	if fg < 0 {
		fg = 39
	}
	if bg < 0 {
		bg = 49
	}

	fg += 30
	bg += 40

	// fmt.Printf("Asked for %d;%d;%d when %s\n", attr, fg, bg, s)
	color += fmt.Sprintf("\033[%d;%d;%dm", attr, fg, bg)
	reset := "\033[0m"
	return color + s + reset
}

// RGB returns the color code corresponding a RGB value set. Arguments
// must range from 0 to 5. Returned values range from 16 (black) to
// 231 (white).
func (t *Legacy) RGB(red, green, blue int) int {
	// handle gray as special case
	if red == green && green == blue && red != 0 {
		return 8
	}

	red = convertColorBit(red)
	green = convertColorBit(green)
	blue = convertColorBit(blue)

	bold := 0
	var ok bool
	if red, ok = checkBoldBit(red); ok {
		bold = 1
	}
	if blue, ok = checkBoldBit(blue); ok {
		bold = 1
	}
	if green, ok = checkBoldBit(green); ok {
		bold = 1
	}

	return red + (green * 2) + (blue * 4) + (bold * 8)
}

func (t *Legacy) AssignColor(s string) int {
	fg := stringId(s, 16)

	// prevent black, white, gray
	if fg == 0 {
		fg += 1
	}
	if fg == 8 {
		fg += 1
	}
	if fg == 7 || fg == 15 {
		fg -= 1
	}

	return fg
}

func convertColorBit(bit int) int {
	switch {
	case bit < 2:
		return 0
	case bit < 4:
		return 1
	case bit < 6:
		return 2
	}
	fmt.Printf("invalid bit: %v\n", bit)
	panic("unreachable")
}

func checkBoldBit(bit int) (int, bool) {
	switch {
	case bit < 2:
		return bit, false
	case bit == 2:
		return 1, true
	}
	fmt.Printf("invalid bit: %v\n", bit)
	panic("unreachable")
}

func checkBold(code int) (int, bool) {
	switch {
	case code < 8:
		return code, false
	case code < 16:
		return code - 8, true
	}
	fmt.Printf("invalid code: %v\n", code)
	panic("unreachable")
}
