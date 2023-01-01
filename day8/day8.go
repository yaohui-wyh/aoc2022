package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	data   [][]int
	result [][]bool
	score  [][]Score
)

type Score struct {
	left, right, top, bottom int
}

func (s Score) calcScore() int {
	return s.left * s.right * s.top * s.bottom
}

func init() {
	bytes, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(bytes), "\n")

	rowSize, colSize := len(lines), len(lines[0])
	data = make([][]int, rowSize)
	result = make([][]bool, rowSize)
	score = make([][]Score, rowSize)
	for i, line := range lines {
		data[i] = make([]int, colSize)
		result[i] = make([]bool, colSize)
		score[i] = make([]Score, colSize)
		for j, char := range line {
			data[i][j] = int(char - '0')
		}
	}
}

func main() {
	calcVisibility()
	calcScore()

	// puzzle1
	count := 0
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[0]); j++ {
			if result[i][j] {
				count++
			}
		}
	}
	fmt.Println(count)

	// puzzle2
	maxScore := 0
	for i := 0; i < len(score); i++ {
		for j := 0; j < len(score[0]); j++ {
			s := score[i][j].calcScore()
			if s > maxScore {
				maxScore = s
			}
		}
	}
	fmt.Println(maxScore)
}

// Calculate visibility from each direction (left -> right, etc.)
// for each direction, if the current item's value is greater than the highest value
// it is visible from that direction
func calcVisibility() {
	rowSize, colSize := len(data), len(data[0])
	// save 1 pass by iterating from left -> right and right -> left at the same time
	for i := 0; i < colSize; i++ {
		highestLeft, highestRight := -1, -1
		for j := 0; j < rowSize; j++ {
			// from left to right comparison
			if data[i][j] > highestLeft {
				result[i][j] = true
				highestLeft = data[i][j]
			}
			// from right to left comparison
			jRight := colSize - 1 - j
			if data[i][jRight] > highestRight {
				result[i][jRight] = true
				highestRight = data[i][jRight]
			}
		}
	}
	for j := 0; j < rowSize; j++ {
		highestTop, highestBottom := -1, -1
		for i := 0; i < colSize; i++ {
			// from top to bottom comparison
			if data[i][j] > highestTop {
				result[i][j] = true
				highestTop = data[i][j]
			}
			// from bottom to top comparison
			iBottom := rowSize - 1 - i
			if data[iBottom][j] > highestBottom {
				result[iBottom][j] = true
				highestBottom = data[iBottom][j]
			}
		}
	}
}

// Calculate score from each direction (left -> right, etc.)
// for each item, calculate the distance between it and the highest value from that direction
func calcScore() {
	rowSize, colSize := len(data), len(data[0])
	for i := 0; i < rowSize; i++ {
		for j := 0; j < colSize; j++ {
			// from left to right
			for k := j - 1; k >= 0; k-- {
				if data[i][j] <= data[i][k] {
					score[i][j].left = j - k
					break
				}
				// current item is the highest one from left -> right
				if k == 0 {
					score[i][j].left = j
				}
			}
			// from right to left
			jRight := colSize - 1 - j
			for k := jRight + 1; k < colSize; k++ {
				if data[i][jRight] <= data[i][k] {
					score[i][jRight].right = k - jRight
					break
				}
				if k == colSize-1 {
					score[i][jRight].right = colSize - 1 - jRight
				}
			}
		}
	}
	for j := 0; j < colSize; j++ {
		for i := 0; i < rowSize; i++ {
			// from top to bottom
			for k := i - 1; k >= 0; k-- {
				if data[i][j] <= data[k][j] {
					score[i][j].top = i - k
					break
				}
				if k == 0 {
					score[i][j].top = i
				}
			}
			// from bottom to top
			iBottom := rowSize - 1 - i
			for k := iBottom + 1; k < rowSize; k++ {
				if data[iBottom][j] <= data[k][j] {
					score[iBottom][j].bottom = k - iBottom
					break
				}
				if k == rowSize-1 {
					score[iBottom][j].bottom = rowSize - 1 - iBottom
				}
			}
		}
	}
}