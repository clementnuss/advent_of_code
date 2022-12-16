package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	x, y int
}

func parsePos(s string) Pos {
	spl := strings.Split(s, ",")
	x, _ := strconv.Atoi(spl[0])
	y, _ := strconv.Atoi(spl[1])
	return Pos{x, y}
}

type Block string

const (
	Sand Block = "o"
	Wall       = "#"
	Air        = "."
)

func main() {

	log.Println("AOC - 2022.12.14")

	inputBytes, err := os.ReadFile("../input")
	// inputBytes, err := os.ReadFile("../test_input")
	if err != nil {
		log.Fatalf("couldn't open input file. error %v", err)
		return
	}

	input := strings.Split(string(inputBytes), "\n")

	cave := make(map[Pos]Block)
	lowestWalls := make(map[int]int)

	for _, line := range input {
		if line == "" {
			continue
		}

		pathPoints := strings.Split(line, " -> ")
		path := make([]Pos, len(pathPoints))

		for i, point := range pathPoints {
			currPos := parsePos(point)
			path[i] = currPos
			cave[currPos] = Wall

			if i != 0 {
				wall := makePath(path[i-1], path[i])
				for _, pos := range wall {
					cave[pos] = Wall
				}
			}
		}
	}

	absoluteLowest := 0

	for pos := range cave {
		if lowest, ok := lowestWalls[pos.x]; ok {
			if pos.y < lowest {
				lowestWalls[pos.x] = pos.y
			}
		} else {
			lowestWalls[pos.x] = pos.y
		}

		if pos.y > absoluteLowest {
			absoluteLowest = pos.y
		}
	}

	for x := -500; x < 1500; x++ {
		cave[Pos{x, absoluteLowest + 2}] = Wall
	}

	for sand := 0; ; sand++ {
		s := Pos{500, 0}

		fell := false

		for {
			if _, ok := cave[Pos{s.x, s.y + 1}]; !ok { // nothing down
				s.y++
			} else if _, ok := cave[Pos{s.x - 1, s.y + 1}]; !ok { // nothing diagonal left
				s.x--
				s.y++
			} else if _, ok := cave[Pos{s.x + 1, s.y + 1}]; !ok { // nothing diagonal right
				s.x++
				s.y++
			} else { // sand block has come to rest
				if s.x == 500 && s.y == 0 {
					fmt.Println("sand source blocked after block: ", sand+1)
					fell = true
					break
				}
				cave[s] = Sand
				break
			}
		}

		if fell || sand > 1e5 {
			break
		}
	}
}

func sgn(x int) int {
	if x < 0 {
		return -1
	} else {
		return 1
	}
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	} else {
		return x
	}
}

func makePath(p1, p2 Pos) []Pos {
	deltaX := p2.x - p1.x
	deltaY := p2.y - p1.y
	if deltaX != 0 {
		path := make([]Pos, 0, abs(deltaX))

		for x := p1.x; x != p2.x; x += sgn(deltaX) {
			path = append(path, Pos{x, p1.y})
		}

		return path
	} else if deltaY != 0 {
		path := make([]Pos, 0, abs(deltaY))

		for y := p1.y; y != p2.y; y += sgn(deltaY) {
			path = append(path, Pos{p1.x, y})
		}

		return path
	}

	return make([]Pos, 0)
}
