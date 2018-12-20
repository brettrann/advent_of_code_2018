package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"io/ioutil"
)

// part1 strategy
type guard struct {
	id int
	minute_seen map[int]int
	total_slept int
}

// find guard with most minutes slept.
// find the minute that guard is most likely to be sleeping
// all sleep is in the same hour, so we just care about minutes
func part1(records []string) (int) {
	guards := buildGuardArray(records)

	// find the guard that slept the most
	var sleepiest = guard{}
	for _, guard := range guards {
		if sleepiest.total_slept < guard.total_slept {
			sleepiest = guard
		}
	}
	// find the most slept minute
	var max int
	var minute int
	for k, v := range sleepiest.minute_seen {
		if max < v {
			max = v
			minute = k
		}
	}
	return sleepiest.id * minute
}

func part2(records []string) (int) {
	guards := buildGuardArray(records)
	var outer_max int
	var outer_minute int
	var max_guard int
	for _, guard := range guards {
		// find the most slept minute
		var max int
		var minute int
		for k, v := range guard.minute_seen {
			if max < v {
				max = v
				minute = k
			}
		}
		if outer_max < max {
			max_guard = guard.id
			outer_max = max
			outer_minute = minute
		}

	}
	//fmt.Printf("%v, %v, %v\n", max_guard, outer_max, outer_minute)
	return max_guard * outer_minute
}

func buildGuardArray(records []string) (map[int]guard) {
	sort.Strings(records)
	var guards = map[int]guard{}
	var id int
	var last_slept int
	for _, line := range records {
		var y, m, d, hr, mn int
		_, err := fmt.Sscanf(line, "[%d-%d-%d %d:%d]", &y, &m, &d, &hr, &mn)
		// we have a blank line that ends up at the top. w/e
		if err != nil {
			continue
		}
		line = line[19:]
		if strings.Contains(line, "begins shift") {
			fmt.Sscanf(line, "Guard #%d begins shift", &id)
			g, ok := guards[id]
			if ok == false {
				g = guard{id: id}
				g.minute_seen = make(map[int]int)
			  guards[g.id] = g
			}
		} else {
			g := guards[id]
			if strings.Contains(line, "falls asleep") {
				last_slept = mn
			} else if strings.Contains(line, "wakes up") {
				g.total_slept += mn - last_slept
				guards[g.id] = g // surprised me. without it total_slept is copied?
				for i := last_slept; i < mn; i++ {
				g.minute_seen[i]++
				}
			}
		}
	}
	return guards
}


func in(path string) ([]string) {
	content, _ := ioutil.ReadFile(path)
	lines := strings.Split(string(content), "\n")
	return lines
}

func main() {
	records := in("input.txt")
	log.Printf("part1: %d", part1(records))
	log.Printf("part2: %d", part2(records))
}

