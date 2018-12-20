package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)


func part1(polymer string) (int) {
	s := []rune(polymer)
	for i := 0; i < len(s)-1; i++ {
		a := strings.ToUpper(string(s[i]))
		b := strings.ToUpper(string(s[i+1]))
		if a == b && s[i] != s[i+1] {
			s = append(s[:i], s[i+2:]...)
			i = -1
			continue
		}
	}
	return len(s)
}

func part2(polymer string) (int) {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	min := len(polymer)
	for _, c := range alphabet {
		attempt := strings.Replace(polymer, string(c), "", -1)
		attempt = strings.Replace(attempt, strings.ToUpper(string(c)), "", -1)
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
