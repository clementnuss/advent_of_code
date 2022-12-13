package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
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
	} else {
		deltaY = sgn(deltaY)
	}
	t.x += deltaX
	t.y += deltaY
}

func print(tail, head *Pos) {
	for y := 4; y >= 0; y-- {
		for x := 0; x <= 5; x++ {
			if x == 0 && y == 0 {
				fmt.Print("s")
			} else if x == tail.x && y == tail.y {
				fmt.Print("T")
			} else if x == head.x && y == head.y {
				fmt.Print("H")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func printMap(m map[Pos]interface{}) {
	for y := 4; y >= 0; y-- {
		for x := 0; x <= 5; x++ {
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

	head, tail := Pos{0, 0}, Pos{0, 0}
	tailMap := make(map[Pos]interface{})

	if _, ok := tailMap[tail]; !ok {
		tailMap[tail] = nil
	}

	input := strings.Split(string(inputBytes), "\n")
	for _, move := range input {
		spl := strings.Split(move, " ")
		dir := spl[0]
		length, _ := strconv.Atoi(spl[1])
		switch dir {
		case "R":
			for i := 0; i < length; i++ {
				head.x++
				print(&tail, &head)
				tail.followHead(&head)
				print(&tail, &head)
				if _, ok := tailMap[tail]; !ok {
					tailMap[tail] = nil
				}
			}
		case "L":
			for i := 0; i < length; i++ {
				head.x--
				print(&tail, &head)
				tail.followHead(&head)
				print(&tail, &head)
				if _, ok := tailMap[tail]; !ok {
					tailMap[tail] = nil
				}
			}
		case "U":
			for i := 0; i < length; i++ {
				head.y++
				print(&tail, &head)
				tail.followHead(&head)
				print(&tail, &head)
				if _, ok := tailMap[tail]; !ok {
					tailMap[tail] = nil
				}
			}
		case "D":
			for i := 0; i < length; i++ {
				print(&tail, &head)
				head.y--
				tail.followHead(&head)
				print(&tail, &head)
				if _, ok := tailMap[tail]; !ok {
					tailMap[tail] = nil
				}
			}
		}
	}

	printMap(tailMap)

	fmt.Println("total score:", len(tailMap))
}
