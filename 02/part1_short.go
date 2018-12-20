package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	fileHandle, _ := os.Open("input1.txt")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	countTwo := 0
	countThree := 0

	// find occurrences of exactly 2 or 3 of a character
	for fileScanner.Scan() {
		charCounts := make(map[rune] int)
		for _, char := range fileScanner.Text() {
			charCounts[char] += 1
		}
		sawTwo := false
		sawThree := false
		for _, count := range charCounts {
			switch count {
			case 2:
				sawTwo = true
			case 3:
				sawThree = true
			}
		}
		if sawTwo {
			countTwo++
		}
		if sawThree {
			countThree++
		}
	}
	fmt.Println(countTwo*countThree)
}
