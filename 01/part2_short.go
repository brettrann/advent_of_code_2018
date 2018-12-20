package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	frequency := 0
	seen := make(map[int]bool)

	for {
		fileHandle, _ := os.Open("input1.txt")
		defer fileHandle.Close()
		fileScanner := bufio.NewScanner(fileHandle)

		for fileScanner.Scan() {
			deviation, _ := strconv.Atoi(fileScanner.Text())
			frequency += deviation
			_, ok := seen[frequency]
			if ok {
				fmt.Println(frequency)
				return
			} else {
				seen[frequency] = true
			}
		}
	}
}
