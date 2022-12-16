package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content, _ := os.ReadFile("day3/input.txt")
	lines := strings.Split(string(content), "\n")
	fmt.Println(puzzle1(lines))
	fmt.Println(puzzle2(lines))
}

func findDuplicateScore(sets ...map[int]bool) int {
	for _, s := range sets {
		for k := range s {
			count := 0
			for _, s2 := range sets {
				if s2[k] {
					count += 1
				}
			}
			if count == len(sets) {
				return k
			}
		}
	}
	return 0
}

func toSet(line string) map[int]bool {
	s := make(map[int]bool)
	for _, c := range line {
		s[convertToInt(c)] = true
	}
	return s
}

func convertToInt(c rune) int {
	if c >= 'a' && c <= 'z' {
		return int(c-'a') + 1
	} else {
		return int(c-'A') + 27
	}
}

func puzzle1(lines []string) int {
	var sum = 0
	for _, line := range lines {
		l1 := line[:len(line)/2]
		l2 := line[len(line)/2:]
		sum += findDuplicateScore(toSet(l1), toSet(l2))
	}
	return sum
}

func puzzle2(lines []string) int {
	var sum = 0
	for i := 0; i < len(lines); i += 3 {
		s := lines[i : i+3]
		sum += findDuplicateScore(toSet(s[0]), toSet(s[1]), toSet(s[2]))
	}
	return sum
}