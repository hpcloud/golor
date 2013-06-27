// 256 color terminal printing in Go.
package golor

// Copyright (c) 2013 ActiveState Software Inc. All rights reserved.

import (
	"crypto/sha1"
	"os"
)

// 16-colors
var RED, GREEN, BLUE, CYAN, MAGENTA, YELLOW, BLACK, WHITE, GRAY int
var R, G, C, M, Y, W int
var MIN, MAX int

var TERM Terminal

type Terminal interface {
	// RGB returns the color code corresponding a RGB value set.
	RGB(red, green, blue int) int
	Colorize(s string, fg int, bg int) string
	// Assign and return an unique color for this string. Automatically
	// avoid black/white colors.
	AssignColor(s string) int
}

func RGB(red, green, blue int) int {
	return TERM.RGB(red, green, blue)
}

func Colorize(s string, fg int, bg int) string {
	return TERM.Colorize(s, fg, bg)
}

func AssignColor(s string) int {
	return TERM.AssignColor(s)
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
	switch os.Getenv("TERM") {
	case "xterm-256color":
		TERM = &Xterm256{}
	default:
		TERM = &Legacy{}
	}

	RED = RGB(5, 0, 0)
	GREEN = RGB(0, 5, 0)
	BLUE = RGB(0, 0, 5)
	CYAN = RGB(0, 5, 5)
	GRAY = RGB(2, 2, 2)
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

	MIN = RGB(0, 0, 0)
	MAX = RGB(5, 5, 5)
}
