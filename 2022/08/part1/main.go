package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Map struct {
	width, height int
	hM            [][][]int
}

type Direction int

const (
	Self Direction = iota
	Up
	Down
	Left
	Right
)

type Pos struct {
	i, j int
}

func (m *Map) initializeMap(input []string) {
	m.width = len(input[0])
	m.height = len(input)

	m.hM = make([][][]int, Right+1)

	for direction := Self; direction <= Right; direction++ {
		m.hM[direction] = make([][]int, m.height)
		for j := 0; j < m.height; j++ {
			m.hM[direction][j] = make([]int, m.width)
		}
	}

	for i, line := range input {
		for j, height := range line {
			m.hM[Self][i][j], _ = strconv.Atoi(string(height))
		}
	}

	for i := 0; i < m.width; i++ {
		for j := 0; j < m.width; j++ {
			for _, dir := range []Direction{Up, Left} {
				m.hM[dir][i][j] = m.checkDir(dir, i, j)
			}
		}
	}

	for i := m.height - 1; i >= 0; i-- {
		for j := m.width - 1; j >= 0; j-- {
			for _, dir := range []Direction{Down, Right} {
				m.hM[dir][i][j] = m.checkDir(dir, i, j)
			}
		}
	}
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func (m *Map) checkDir(dir Direction, i, j int) int {
	self := m.hM[Self][i][j]
	switch dir {
	case Up:
		return Max(m.get(dir, i-1, j), self)
	case Down:
		return Max(m.get(dir, i+1, j), self)
	case Left:
		return Max(m.get(dir, i, j-1), self)
	case Right:
		return Max(m.get(dir, i, j+1), self)
	}

	panic("shouldn't be here")
}

func (m *Map) get(dir Direction, i, j int) int {
	if i < 0 || i >= m.height || j < 0 || j >= m.width {
		return -1
	}

	return m.hM[dir][i][j]
}

func (m *Map) getPos(dir Direction, p Pos) int {
	i, j := p.i, p.j
	if i < 0 || i >= m.height || j < 0 || j >= m.width {
		return -1
	}

	return m.hM[dir][i][j]
}

func index(dir Direction, p Pos) Pos {
	switch dir {
	case Up:
		p.i++
	case Down:
		p.i--
	case Right:
		p.j++
	case Left:
		p.j--
	}
	return p
}

func (m *Map) countVisible() int {
	total := 0
	for i := 0; i < m.height; i++ {
		for j := 0; j < m.width; j++ {
			self := m.hM[Self][i][j]
			if self > m.get(Up, i-1, j) ||
				self > m.get(Down, i+1, j) ||
				self > m.get(Right, i, j+1) ||
				self > m.get(Left, i, j-1) {
				total++
			}
		}
	}

	return total
}

func main() {

	log.Println("AOC - 2022.12.08+1")

	inputBytes, err := os.ReadFile("../input")
	// inputBytes, err := os.ReadFile("../test_input")
	if err != nil {
		log.Fatalf("couldn't open input file. error %v", err)
		return
	}

	input := strings.Split(string(inputBytes), "\n")

	var m = Map{}
	m.initializeMap(input)

	fmt.Println("total score:", m.countVisible())
}
