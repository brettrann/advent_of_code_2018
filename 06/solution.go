package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)


func part1(polymer string) (int) {
	a := []rune(polymer)
	for i := 0; i < len(a)-1; i++ {
		if strings.ToUpper(string(a[i])) == strings.ToUpper(string(a[i+1])) {
			if a[i] != a[i+1] {
				a = append(a[:i], a[i+2:]...)
				i = -1
				continue
			}
		}
	}
	return len(a)
}

func part2(polymer string) (int) {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	min := len(polymer)
	for _, c := range alphabet {
		attempt := strings.Replace(strings.Replace(polymer, string(c), "", -1), strings.ToUpper(string(c)), "", -1)
		size := part1(attempt)
		if size < min {
			min = size
		}
	}
	return min
}


func in(path string) ([]string) {
	content, _ := ioutil.ReadFile(path)
	lines := strings.Split(string(content), "\n")
	lines = lines[:len(lines)-1]
	return lines
}

func main() {
	polymers := in("input.txt")
	fmt.Printf("part1: %d\n", part1(polymers[0]))
	fmt.Printf("part2: %d\n", part2(polymers[0]))
}
