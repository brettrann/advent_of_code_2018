package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	fileHandle, _ := os.Open("input1.txt")
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	frequency := 0

	for fileScanner.Scan() {
		deviation, _ := strconv.Atoi(fileScanner.Text())
		frequency += deviation
	}

	fmt.Println(frequency)
}
