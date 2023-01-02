package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	bytes, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(bytes), "\n")

	puzzle1(lines)
	puzzle2(lines)
}

func puzzle1(lines []string) {
	register := 1
	sum := 0
	cyclesCh := make(chan int)
	doneCh := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		cycles := 1
		for {
			// increment cycles channel
			cycles++
			cyclesCh <- 1

			// wait for cycle to finish
			<-doneCh

			if (cycles-20)%40 == 0 {
				fmt.Printf("cycles: %d, register: %d\n", cycles, register)
				sum += cycles * register
			}
		}
	}()

	go func() {
		for _, line := range lines {
			if line != "noop" {
				// cycle 1 finished
				<-cyclesCh
				doneCh <- true

				// cycle 2 finished
				<-cyclesCh
				op, _ := strconv.Atoi(line[5:])
				register += op
				doneCh <- true
			} else {
				// noop finished in 1 cycle
				<-cyclesCh
				doneCh <- true
			}
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(sum)
}

func puzzle2(lines []string) {
	cyclesCh := make(chan int)
	doneCh := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(1)

	crtIdx, spiritIdx := 0, 0
	go func() {
		for {
			// print CRT line
			if spiritIdx <= crtIdx && crtIdx <= spiritIdx+2 {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}

			// inc & line break
			crtIdx++
			if crtIdx == 40 {
				fmt.Println()
				crtIdx = 0
			}

			cyclesCh <- 1
			<-doneCh
		}
	}()

	go func() {
		for _, line := range lines {
			if line != "noop" {
				<-cyclesCh
				doneCh <- true

				<-cyclesCh
				op, _ := strconv.Atoi(line[5:])
				spiritIdx += op
				doneCh <- true
			} else {
				<-cyclesCh
				doneCh <- true
			}
		}
		wg.Done()
	}()

	wg.Wait()
}