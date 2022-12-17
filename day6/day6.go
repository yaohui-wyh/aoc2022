package main

import (
	"fmt"
	"os"
)

func main() {
	content, _ := os.ReadFile("input.txt")
	line := string(content)
	fmt.Println(findStart(line, 4))
	fmt.Println(findStart(line, 14))
}

func findStart(line string, size int) int {
	for i := 0; i < len(line); {
		window := map[byte]int{}
		for j := 0; j < size; j++ {
			c := line[i+j]
			idx, ok := window[c]
			if ok {
				i = idx + 1
				break
			}
			window[c] = i + j
			if j == size-1 {
				return i + j + 1
			}
		}
	}
	return -1
}
