package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	win  = 6
	draw = 3
	lost = 0
)

func main() {
	content, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(content), "\n")
	fmt.Println(puzzle1(lines))
	fmt.Println(puzzle2(lines))
}

func parseScore(name string) int {
	switch name {
	case "A", "X":
		return 1
	case "B", "Y":
		return 2
	case "C", "Z":
		return 3
	}
	return 0
}

func calculateScore(a, b int) int {
	if a == b {
		return b + draw
	}
	if (a == 1 && b == 2) || (a == 2 && b == 3) || (a == 3 && b == 1) {
		return b + win
	}
	return b + lost
}

func parseScoreFromResult(a int, encResult string) int {
	var result = -1
	switch encResult {
	case "X":
		result = lost
	case "Y":
		result = draw
	case "Z":
		result = win
	}

	switch result {
	case lost:
		switch a {
		case 1:
			return 3
		case 2:
			return 1
		case 3:
			return 2
		}
	case win:
		switch a {
		case 1:
			return 2
		case 2:
			return 3
		case 3:
			return 1
		}
	case draw:
		return a
	}
	return -1
}

func puzzle1(lines []string) int {
	var totalScore = 0
	for _, line := range lines {
		s := strings.Split(line, " ")
		a := parseScore(s[0])
		b := parseScore(s[1])
		totalScore += calculateScore(a, b)
	}
	return totalScore
}

func puzzle2(lines []string) int {
	var totalScore = 0
	for _, line := range lines {
		s := strings.Split(line, " ")
		a := parseScore(s[0])
		b := parseScoreFromResult(a, s[1])
		totalScore += calculateScore(a, b)
	}
	return totalScore
}
