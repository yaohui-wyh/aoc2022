package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("day4/input.txt")
	lines := strings.Split(string(content), "\n")
	fmt.Println(puzzle1(lines))
	fmt.Println(puzzle2(lines))
}

type Range struct {
	l int
	r int
}

func NewRange(line string) Range {
	pair := strings.Split(line, "-")
	l, _ := strconv.Atoi(pair[0])
	r, _ := strconv.Atoi(pair[1])
	return Range{l, r}
}

func isContained(r1, r2 Range) bool {
	return (r1.l <= r2.l && r1.r >= r2.r) ||
		(r1.l >= r2.l && r1.r <= r2.r)
}

func isOverlapped(r1, r2 Range) bool {
	return (r1.l <= r2.l && r2.l <= r1.r) ||
		(r1.l >= r2.l && r1.l <= r2.r)
}

func puzzle1(lines []string) int {
	var sum = 0
	for _, line := range lines {
		parts := strings.Split(line, ",")
		if isContained(NewRange(parts[0]), NewRange(parts[1])) {
			sum += 1
		}
	}
	return sum
}

func puzzle2(lines []string) int {
	var sum = 0
	for _, line := range lines {
		parts := strings.Split(line, ",")
		if isOverlapped(NewRange(parts[0]), NewRange(parts[1])) {
			sum += 1
		}
	}
	return sum
}