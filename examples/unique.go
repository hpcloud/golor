// Copyright (c) 2013 ActiveState Software Inc. All rights reserved.

package main

import (
	"fmt"
	"github.com/ActiveState/golor"
)

func main() {
	fmt.Println("# Consistently assigning unique colors to strings:")
	for _, ape := range getScientificApes() {
		fmt.Println(golor.Colorize(ape, golor.AssignColor(ape), -1))
	}
}

func getScientificApes() []string {
	// http://en.wikipedia.org/wiki/List_of_individual_apes#Scientific_apes
	return []string{
		"Abang",
		"Bonnie",
		"Chantek",
		"Enos",
		"Frodo",
		"Kanzi",
		"Koko",
		"Lucy",
		"Nyota",
		"Sultan",
		"Titus",
		"Washoe",
	}
}
