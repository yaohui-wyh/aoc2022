package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := os.ReadFile("./input.txt")
	lines := strings.Split(string(bytes), "\n")

	r := NewRope()
	r2 := NewRope2()
	for _, line := range lines {
		direction := line[:1]
		step, _ := strconv.Atoi(line[2:])
		for i := 0; i < step; i++ {
			r.Move(direction)
			r2.Move(direction)
		}
	}
	fmt.Println(len(r.TailVisitedMap))
	fmt.Println(len(r2.knobs[len(r2.knobs)-1].TailVisitedMap))
}

type Position struct {
	X, Y int
}

type Rope struct {
	Head           Position
	Tail           Position
	TailVisitedMap map[Position]bool
}

func NewRope() *Rope {
	r := &Rope{
		Head:           Position{0, 0},
		Tail:           Position{0, 0},
		TailVisitedMap: map[Position]bool{},
	}
	r.TailVisitedMap[r.Tail] = true
	return r
}

func (r *Rope) Move(direction string) {
	r.moveHead(direction)
	r.adjustPosition()
}

func (r *Rope) moveHead(direction string) {
	switch direction {
	case "R":
		r.Head.X++
	case "L":
		r.Head.X--
	case "U":
		r.Head.Y++
	case "D":
		r.Head.Y--
	}
}

// this is how Rope is rebalanced...
func (r *Rope) adjustPosition() {
	dx := r.Head.X - r.Tail.X
	dy := r.Head.Y - r.Tail.Y

    // copied from some smart guys' answer...
	update := false
	if dx == 2 || dx == -2 {
		dx /= 2
		update = true
	}
	if dy == 2 || dy == -2 {
		dy /= 2
		update = true
	}
	if update {
		r.Tail.X += dx
		r.Tail.Y += dy
	}
	r.TailVisitedMap[r.Tail] = true
}

type Rope2 struct {
	knobs []*Rope
}

func NewRope2() *Rope2 {
	r := &Rope2{}
	for i := 0; i < 9; i++ {
		r.knobs = append(r.knobs, NewRope())
	}
	return r
}

func (r *Rope2) Move(direction string) {
	r.knobs[0].Move(direction)
	for i := 1; i < len(r.knobs); i++ {
		// "Move" mannually
		r.knobs[i].Head = r.knobs[i-1].Tail
		r.knobs[i].adjustPosition()
	}
}
