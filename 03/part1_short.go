

package main

import (
	"bufio"
	"fmt"
	"os"
)

// eg: #32 @ 863,904: 22x20
type claim struct {
	id		 int
	x,y		 int
	w,h		 int
}

func main() {
	fileHandle, _ := os.Open("input1.txt")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

// find size of overlapping claim size
// 1000x1000 sheet
// build marker counts based on each claim
// sum squares that are > 1

	var sheet [1000][1000]int

// since we only care about overlaps we could make a map
// of used coords and then count the ones with > 1. w/e.

	overlap := 0
	for fileScanner.Scan() {
		var c claim

		var line = fileScanner.Text()
		fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &c.id, &c.x, &c.y, &c.w, &c.h)

		// mark out the square. count overlaps if 1 layer already there.
		for y := c.y; y < c.y+c.h; y++ {
			for x := c.x; x < c.x+c.w; x++ {
				sheet[x][y]++
				if sheet[x][y] == 2 {
					overlap++
				}
			}
		}
	}

	fmt.Println(overlap)
}
