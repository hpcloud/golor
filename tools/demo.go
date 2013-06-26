package main

import (
	"fmt"
	"logyard/util/golor"
)

func main() {
	fmt.Println(golor.Colorize("GRAY", golor.GRAY, -1))
	fmt.Println(golor.Colorize("RED", golor.R, -1))
	fmt.Println(golor.Colorize("GREEN", golor.G, -1))
	fmt.Println(golor.Colorize("BLUE", golor.BLUE, -1))
	fmt.Println(golor.Colorize("MAGENTA", golor.M, -1))
	fmt.Println(golor.Colorize("YELLOW", golor.Y, -1))
	fmt.Println(golor.Colorize("CYAN", golor.C, -1))
	fmt.Println(golor.Colorize("BLACK", golor.BLACK, golor.WHITE))
	fmt.Println(golor.Colorize("WHITE", golor.WHITE, -1))
	fmt.Println(golor.Colorize("LOWEST", golor.RGB(0, 0, 0), golor.WHITE))
	fmt.Println(golor.Colorize("HIGHEST", golor.RGB(5, 5, 5), -1))
	// fmt.Println(golor.Colorize("TEST", 16, -1))
	/* for i := 255; i >= 0; i-- {
		fmt.Println(golor.Colorize("Foreground", i, -1))
		// fmt.Println(golor.ColorizeBg("Background", i, 100))
	}*/
}
