package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Pos struct {
	x, y int
}

func abs(x int) int {
	if x < 0 {
		x *= -1
	}

	return x
}

func sgn(x int) int {
	if x < 0 {
		return -1
	} else {
		return 1
	}
}

func (t *Pos) dist(h *Pos) float64 {
	deltaX := h.x - t.x
	deltaY := h.y - t.y

	return math.Sqrt(math.Pow(float64(deltaX), 2) + math.Pow(float64(deltaY), 2))
}

func (t *Pos) followHead(h *Pos) {
	dist := t.dist(h)
	if dist <= math.Sqrt2 {
		return
	}

	deltaX := h.x - t.x
	deltaY := h.y - t.y

	if dist == 2 {
		if abs(deltaX) > 0 {
			t.x += sgn(deltaX)
		} else if abs(deltaY) > 0 {
			t.y += sgn(deltaY)
		} else {
			panic("shouldn't be here")
		}

		return
	}

	if abs(deltaX) > abs(deltaY) {
		deltaX = sgn(deltaX)
	} else if abs(deltaY) > abs(deltaX) {
		deltaY = sgn(deltaY)
	} else {
		deltaX = sgn(deltaX)
		deltaY = sgn(deltaY)
	}

	t.x += deltaX
	t.y += deltaY
}

func print(rope []*Pos) {

	m := make(map[Pos]string, len(rope))

	for i, knot := range rope {
		if i == 0 {
			m[*knot] = "H"
			continue
		}

		m[*knot] = fmt.Sprint(i)
	}

	for y := 24; y >= 0; y-- {
		for x := 0; x <= 25; x++ {
			if x == 0 && y == 0 {
				fmt.Print("s")
			} else if s, ok := m[Pos{x, y}]; ok {
				fmt.Print(s)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	fmt.Println()
	time.Sleep(100 * time.Millisecond)
}

func printMap(m map[Pos]interface{}) {
	for y := 24; y >= 0; y-- {
		for x := 0; x <= 25; x++ {
			if x == 0 && y == 0 {
				fmt.Print("s")
			} else if _, ok := m[Pos{x, y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func main() {
	log.Println("AOC - 2022.12.09+1")

	inputBytes, err := os.ReadFile("../input")
	if err != nil {
		log.Fatalf("couldn't open input file. error %v", err)
		return
	}

	rope := make([]*Pos, 10)
	for i := 0; i < len(rope); i++ {
		rope[i] = &Pos{x: 8, y: 4}
	}

	tail := rope[len(rope)-1]
	head := rope[0]

	tailMap := make(map[Pos]interface{})

	input := strings.Split(string(inputBytes), "\n")
	for _, move := range input {
		fmt.Printf("### %v ###\n", move)

		spl := strings.Split(move, " ")
		dir := spl[0]
		length, _ := strconv.Atoi(spl[1])

		for i := 0; i < length; i++ {
			switch dir {
			case "R":
				head.x++
			case "L":
				head.x--
			case "U":
				head.y++
			case "D":
				head.y--
			}

			for i := 1; i < len(rope); i++ {
				t := rope[i]
				t.followHead(rope[i-1])
			}

			if _, ok := tailMap[*tail]; !ok {
				tailMap[*tail] = nil
			}
		}
	}

	printMap(tailMap)

	fmt.Println("total score:", len(tailMap))
}
