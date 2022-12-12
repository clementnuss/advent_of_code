package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Directory struct {
	parent  *Directory
	subdirs map[string]*Directory
	files   map[string]int
	path    string
	size    int
}

func main() {

	log.Println("AOC - 2022.12.07+1")

	inputBytes, err := os.ReadFile("../input")
	if err != nil {
		log.Fatalf("couldn't open input file. error %v", err)
		return
	}

	var rootDir = &Directory{
		parent:  nil,
		subdirs: make(map[string]*Directory),
		path:    "/",
		files:   make(map[string]int),
	}
	var currentDir = rootDir

	input := strings.Split(string(inputBytes), "\n")
	for _, line := range input[1:] {
		spl := strings.Split(line, " ")
		if line[0] == '$' {
			if spl[1] == "cd" && spl[2] == ".." {
				currentDir.parent.size += currentDir.size
				currentDir = currentDir.parent
			} else if spl[1] == "cd" {
				cd := spl[2]
				newDir := Directory{
					parent:  currentDir,
					path:    currentDir.path + cd + "/",
					files:   make(map[string]int),
					subdirs: make(map[string]*Directory),
				}
				if currentDir.subdirs != nil {
					currentDir.subdirs[cd] = &newDir
				}
				currentDir = &newDir
			} else if spl[1] == "ls" {
				continue
			}
		} else if spl[0] == "dir" {
			continue
		} else {
			size, _ := strconv.Atoi(spl[0])
			name := spl[1]
			currentDir.files[name] = size
			currentDir.size += size
		}
	}

	for currentDir != rootDir {
		currentDir.parent.size += currentDir.size
		currentDir = currentDir.parent

	}

	fmt.Println("sum:", rootDir.sumLess100000())
}

func (d *Directory) sumLess100000() (sum int) {
	for _, dir := range d.subdirs {
		sum += dir.sumLess100000()
	}
	if d.size < 1e5 {
		sum += d.size
	}

	return sum
}
