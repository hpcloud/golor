package main

import (
	"fmt"
	"logyard/util/golor"
)

func main() {
	for i := 255; i >= 0; i-- {
		fmt.Println(golor.Colorize("Foreground", i, -1))
		// fmt.Println(golor.ColorizeBg("Background", i, 100))
	}
}
