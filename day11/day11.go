package main

import (
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	Id              int
	Items           []int
	Operation       func(op int) int
	TestDivisor     int
	NextIdIfTrue    int
	NextIdIfFalse   int
	InspectionTimes int
}

func parseOperation(line string) func(op int) int {
	re := regexp.MustCompile("(old) ([+\\-*/]) (old|\\d+)")
	match := re.FindAllStringSubmatch(line, -1)[0]
	operator := match[2]
	operand := match[3]
	switch operator {
	case "+":
		return func(op int) int {
			if operand == "old" {
				return op + op
			} else {
				return op + MustAtoi(operand)
			}
		}
	case "-":
		return func(op int) int {
			if operand == "old" {
				return 0
			} else {
				return op - MustAtoi(operand)
			}
		}
	case "*":
		return func(op int) int {
			if operand == "old" {
				return op * op
			} else {
				return op * MustAtoi(operand)
			}
		}
	case "/":
		return func(op int) int {
			if operand == "old" {
				return 1
			} else {
				// ignore divide by zero
				return op / MustAtoi(operand)
			}
		}
	}
	return nil
}

func StrListToIntList(strList []string) []int {
	intList := make([]int, 0)
	for _, i := range strList {
		item, _ := strconv.Atoi(i)
		intList = append(intList, item)
	}
	return intList
}

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func NewMonkey(lines []string) Monkey {
	monkey := Monkey{}
	re := regexp.MustCompile("(\\d+)")
	monkey.Id = MustAtoi(re.FindString(lines[0]))
	monkey.Items = StrListToIntList(re.FindAllString(lines[1], -1))
	monkey.Operation = parseOperation(lines[2])
	monkey.TestDivisor = MustAtoi(re.FindString(lines[3]))
	monkey.NextIdIfTrue = MustAtoi(re.FindString(lines[4]))
	monkey.NextIdIfFalse = MustAtoi(re.FindString(lines[5]))
	monkey.InspectionTimes = 0
	return monkey
}

// FIXME: overflow for puzzle2
func Process(monkeys []Monkey, idx int) {
	m := monkeys[idx]
	for _, item := range m.Items {
		result := m.Operation(item)
		result /= 3
		if result%m.TestDivisor == 0 {
			monkeys[m.NextIdIfTrue].Items = append(monkeys[m.NextIdIfTrue].Items, result)
		} else {
			monkeys[m.NextIdIfFalse].Items = append(monkeys[m.NextIdIfFalse].Items, result)
		}
	}
	monkeys[idx].InspectionTimes += len(m.Items)
	monkeys[idx].Items = []int{}
}

func main() {
	bytes, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(bytes), "\n")

	var monkeys []Monkey
	for i := 0; i < len(lines); i += 7 {
		instructions := lines[i : i+6]
		monkeys = append(monkeys, NewMonkey(instructions))
	}

	round := 20
	for i := 0; i < round; i++ {
		for idx := range monkeys {
			Process(monkeys, idx)
		}
	}

	var times []int
	for _, m := range monkeys {
		times = append(times, m.InspectionTimes)
	}
	sort.Ints(times)
	println(times[len(times)-1] * times[len(times)-2])
}
