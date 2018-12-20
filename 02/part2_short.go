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

	// compare all possible pairs of lines.
	// iterate and build list and compare with list does it
	// find line that differ by 1 character
	var lines = []string{}
	for fileScanner.Scan() {
		var line = fileScanner.Text()
		for _, otherline := range lines {
			count := 0
			// count matching characters
			for i, _ := range line {
				if line[i] == otherline[i] {
					count++
				}
				// print matching chars if only 1 differs
				if (count == len(line)-1) {
					for i, c := range line {
						if line[i] == otherline[i] {
							fmt.Print(string(c))
						}
					}
				}
			}
		}
		lines = append(lines, line)
	}
	fmt.Println()
}
