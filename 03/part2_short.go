package main

import (
	"bufio"
	"fmt"
	"os"
)

// eg: #32 @ 863,904: 22x20
type claim struct {
	id int
	x,y int
	w,h int
}

func main() {
	fileHandle, _ := os.Open("input1.txt")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)


claims := []claim{}
	var sheet [1000][1000]int

	overlap := 0
	for fileScanner.Scan() {
		var c claim

		var line = fileScanner.Text()
		fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &c.id, &c.x, &c.y, &c.w, &c.h)
		claims = append(claims, c)

		for y := c.y; y < c.y+c.h; y++ {
			for x := c.x; x < c.x+c.w; x++ {
				sheet[x][y]++
				if sheet[x][y] == 2 {
					overlap++
				}
			}
		}
	}

	// part 1
	fmt.Println(overlap)

	//part2
claims_loop:
	for _, c := range claims {
		for y := c.y; y < c.y+c.h; y++ {
			for x := c.x; x < c.x+c.w; x++ {
				if sheet[x][y] > 1 {
					continue claims_loop
				}
			}
		}
		fmt.Println(c)
	}
}
