package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var store = [][]string{
	{"B", "Q", "C"},
	{"R", "Q", "W", "Z"},
	{"B", "M", "R", "L", "V"},
	{"C", "Z", "H", "V", "T", "W"},
	{"D", "Z", "H", "B", "N", "V", "G"},
	{"H", "N", "P", "C", "J", "F", "V", "Q"},
	{"D", "G", "T", "R", "W", "Z", "S"},
	{"C", "G", "M", "N", "B", "W", "Z", "P"},
	{"N", "J", "B", "M", "W", "Q", "F", "P"},
}

type instruction struct {
	count int
	// from is zero-based
	from int
	// to is zero-based
	to int
}

func deepCopy(src [][]string) [][]string {
	dst := make([][]string, len(src))
	for i := range src {
		dst[i] = make([]string, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}

func main() {
	// puzzle1
	s1 := deepCopy(store)
	for _, i := range parseInstructions() {
		move(s1, i)
	}
	result := []string{}
	for _, s := range s1 {
		result = append(result, s[len(s)-1])
	}
	fmt.Println(strings.Join(result, ""))

	// puzzle2
	s2 := deepCopy(store)
	for _, i := range parseInstructions() {
		move2(s2, i)
	}
	result = []string{}
	for _, s := range s2 {
		result = append(result, s[len(s)-1])
	}
	fmt.Println(strings.Join(result, ""))
}

func parseInstructions() []instruction {
	instructions := []instruction{}
	content, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(content), "\n")
	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

	for _, l := range lines {
		matches := re.FindStringSubmatch(l)
		if len(matches) == 4 {
			c, _ := strconv.Atoi(matches[1])
			f, _ := strconv.Atoi(matches[2])
			t, _ := strconv.Atoi(matches[3])
			instructions = append(instructions, instruction{count: c, from: f - 1, to: t - 1})
		}
	}
	return instructions
}

func move(store [][]string, inst instruction) {
	moved := store[inst.from][len(store[inst.from])-inst.count : len(store[inst.from])]
	for i := len(moved) - 1; i >= 0; i-- {
		store[inst.to] = append(store[inst.to], moved[i])
	}
	store[inst.from] = store[inst.from][:len(store[inst.from])-inst.count]
}

func move2(store [][]string, inst instruction) {
	moved := store[inst.from][len(store[inst.from])-inst.count : len(store[inst.from])]
	store[inst.to] = append(store[inst.to], moved...)
	store[inst.from] = store[inst.from][:len(store[inst.from])-inst.count]
}
