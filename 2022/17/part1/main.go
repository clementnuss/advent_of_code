package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Dir int

const (
	Left Dir = iota
	Right
	Down
)

type Pos struct {
	x, y int
}

func (p *Pos) move(d Dir) Pos {
	switch d {
	case Left:
		return Pos{p.x - 1, p.y}
	case Right:
		return Pos{p.x + 1, p.y}
	case Down:
		return Pos{p.x, p.y - 1}
	}
	panic("shouldn't be here")
}

type rock struct {
	blocks []Pos
}

func (r *rock) move(d Dir, chamber map[Pos]string) (bool, rock) {
	newRock := rock{}

	for _, p := range r.blocks {
		newPos := p.move(d)
		if newPos.x == 0 || newPos.x == 8 {
			return false, *r
		}

		if _, ok := chamber[newPos]; ok { // collision!
			return false, *r
		}

		newRock.blocks = append(newRock.blocks, newPos)
	}

	return true, newRock
}

func main() {

	log.Println("AOC - 2022.12.17")

	inputBytes, err := os.ReadFile("../input")
	// inputBytes, err := os.ReadFile("../input")

	if err != nil {
		log.Fatalf("couldn't open input file. error %v", err)
		return
	}

	input := strings.Split(string(inputBytes), "\n")
	jetPattern := make([]Dir, 0, len(input[0]))

	for _, dir := range input[0] {
		switch dir {
		case '<':
			jetPattern = append(jetPattern, Left)
		case '>':
			jetPattern = append(jetPattern, Right)
		}
	}

	currentHighest := 0
	verticalChamber := make(map[Pos]string, 0)

	for x := 1; x < 8; x++ {
		p := Pos{x, currentHighest}
		verticalChamber[p] = "-"
	}

	jetIdx := 0

	for i := 0; i < 2022; i++ {
		var r rock

		switch i % 5 {
		case 0: // --- rock
			for i := 0; i < 4; i++ {
				p := Pos{3 + i, currentHighest + 4}
				r.blocks = append(r.blocks, p)
			}
		case 1: // +  rock
			lowest := currentHighest + 4
			r.blocks = append(r.blocks,
				Pos{3, lowest + 1},
				Pos{4, lowest},
				Pos{4, lowest + 1},
				Pos{4, lowest + 2},
				Pos{5, lowest + 1},
			)
		case 2: //
			lowest := currentHighest + 4
			r.blocks = append(r.blocks,
				Pos{3, lowest},
				Pos{4, lowest},
				Pos{5, lowest},
				Pos{5, lowest + 1},
				Pos{5, lowest + 2},
			)
		case 3: //
			lowest := currentHighest + 4
			r.blocks = append(r.blocks,
				Pos{3, lowest},
				Pos{3, lowest + 1},
				Pos{3, lowest + 2},
				Pos{3, lowest + 3},
			)
		case 4: //
			lowest := currentHighest + 4
			r.blocks = append(r.blocks,
				Pos{3, lowest},
				Pos{3, lowest + 1},
				Pos{4, lowest},
				Pos{4, lowest + 1},
			)
		}

	block:
		for {

			// jet move
			dir := jetPattern[jetIdx%len(jetPattern)]
			jetIdx++
			ok, newRock := r.move(dir, verticalChamber)

			if ok {
				r = newRock
			}

			// fmt.Println("")
			// debugChamber(verticalChamber, 8, r)

			// gravity move
			ok, newRock = r.move(Down, verticalChamber)
			if ok {
				r = newRock
			} else {
				for _, p := range r.blocks {
					verticalChamber[p] = "#"
					if p.y > currentHighest {
						currentHighest = p.y
					}
				}
				break block
			}
		}
		// fmt.Println("block #", i)
		// debugChamber(verticalChamber, 15, rock{})
	}

	fmt.Println("current highest:", currentHighest)
}

func debugChamber(verticalChamber map[Pos]string, currentMax int, rock rock) {
	for y := currentMax; y >= 0; y-- {
	xLoop:
		for x := 0; x < 9; x++ {
			if x == 0 || x == 8 {
				if y == 0 {
					fmt.Print("+")
				} else {
					fmt.Print("|")
				}

				continue
			}

			p := Pos{x, y}
			for _, b := range rock.blocks {
				if p == b {
					fmt.Print("@")
					continue xLoop
				}
			}

			if s, ok := verticalChamber[p]; ok {
				fmt.Print(s)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
