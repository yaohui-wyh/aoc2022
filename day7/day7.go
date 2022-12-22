package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var current *File

func main() {
	dummy := &File{name: "dummy", isDir: true, parent: nil, children: make(map[string]*File)}
	current = dummy

	content, _ := os.ReadFile("./input.txt")
	parseContent(string(content))

	sum := 0
	smallest := 30000000
	total := dummy.children["/"].calculateSize()

	// BFS
	current = dummy
	var queue []*File
	for _, d := range current.children {
		queue = append(queue, d)
	}
	for len(queue) > 0 {
		current = queue[0]
		queue = queue[1:]
		size := current.calculateSize()
		// puzzle1
		if current.isDir && size < 100000 {
			sum += size
		}
		// puzzle2
		if current.isDir && total-size <= 40000000 && size < smallest {
			smallest = size
		}
		for _, d := range current.children {
			queue = append(queue, d)
		}
	}
	fmt.Println("puzzle1", sum)
	fmt.Println("puzzle2", smallest)
}

type File struct {
	name     string
	size     int
	isDir    bool
	children map[string]*File
	parent   *File
}

func parseContent(lines string) {
	for _, line := range strings.Split(lines, "\n") {
		switch {
		case strings.HasPrefix(line, "$ cd "):
			current = current.changeDirectory(line[5:])
		case strings.HasPrefix(line, "$ ls") || strings.HasPrefix(line, "dir "):
			continue
		default:
			// pattern: "<size> <name>"
			parts := strings.Split(line, " ")
			size, _ := strconv.Atoi(parts[0])
			current.children[parts[1]] = &File{name: parts[1], size: size, parent: current, children: make(map[string]*File)}
		}
	}
}

func (f *File) calculateSize() int {
	if !f.isDir {
		return f.size
	}
	sum := 0
	for _, f := range f.children {
		sum += f.calculateSize()
	}
	return sum
}

func (f *File) changeDirectory(name string) *File {
	if name == ".." {
		return f.parent
	} else {
		newFile := &File{name: name, parent: f, isDir: true, children: make(map[string]*File)}
		f.children[name] = newFile
		return newFile
	}
}
