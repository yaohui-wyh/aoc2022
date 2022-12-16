package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(content), "\n")
	fmt.Println(puzzle1(lines))
	fmt.Println(puzzle2(lines))
}

func puzzle1(lines []string) int {
	var maxCalories = 0
	var currentCalories = 0
	for _, line := range lines {
		if line != "" {
			n, _ := strconv.Atoi(line)
			currentCalories += n
		} else {
			if currentCalories > maxCalories {
				maxCalories = currentCalories
			}
			currentCalories = 0
		}
	}
	return maxCalories
}

func puzzle2(lines []string) int {
	var maxCalories = []int{0, 0, 0}
	var currentCalories = 0
	for _, line := range lines {
		if line != "" {
			n, _ := strconv.Atoi(line)
			currentCalories += n
		} else {
			sort.Ints(maxCalories)
			if currentCalories > maxCalories[0] {
				maxCalories[0] = currentCalories
			}
			currentCalories = 0
		}
	}
	return maxCalories[0] + maxCalories[1] + maxCalories[2]
}
