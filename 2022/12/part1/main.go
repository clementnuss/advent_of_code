package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type Pos struct {
	x, y int
}

type Vertex struct {
	p     Pos
	neigh Neighbors
	prev  *Vertex
}

type Neighbors []Pos

var (
	mapHeight int
	mapWidth  int
)

func main() {

	log.Println("AOC - 2022.12.12+1")

	inputBytes, err := os.ReadFile("../input")
	// inputBytes, err := os.ReadFile("../test_input")
	if err != nil {
		log.Fatalf("couldn't open input file. error %v", err)
		return
	}

	input := strings.Split(string(inputBytes), "\n")

	mapWidth = len(input[0])
	mapHeight = len(input)

	var startPos, targetPos Pos

	vertices := make(map[Pos]*Vertex, mapHeight*mapWidth)
	toVisit := make(map[Pos]Neighbors, mapHeight*mapWidth)
	dist := make(map[Pos]int, mapHeight*mapWidth)

	heights := make(map[Pos]int, mapHeight*mapWidth)

	for y, line := range input {
		for x := 0; x < len(line); x++ {
			p := Pos{x: x, y: y}
			switch c := line[x]; c {
			case 'S':
				heights[p] = 0
				dist[p] = 0
				startPos = p

				continue
			case 'E':
				targetPos = p
				heights[p] = int('z' - 'a')
				dist[p] = math.MaxInt
			default:
				heights[p] = int(c - 'a')
				dist[p] = math.MaxInt
			}
		}
	}

	findNeighbors(heights, startPos)

	for p := range heights {
		neigh := findNeighbors(heights, p)

		v := Vertex{
			p:     p,
			neigh: neigh,
		}
		vertices[p] = &v
		toVisit[p] = neigh
	}

	for len(toVisit) > 0 {
		var minPos Pos

		currentMin := math.MaxInt

		for p := range toVisit {
			if dist[p] < currentMin {
				minPos = p
				currentMin = dist[p]
			}
		}

		if currentMin == math.MaxInt {
			fmt.Println("non reachable node")
			break
		}

		neigh := toVisit[minPos]

		delete(toVisit, minPos)

		for _, n := range neigh {
			if _, ok := toVisit[n]; ok {
				if currentMin+1 < dist[n] {
					dist[n] = currentMin + 1
					vertices[n].prev = vertices[minPos]
				}
			}
		}

		// fmt.Println("minPos: ", minPos)
	}

	path := make(map[Pos]string)
	path[targetPos] = "E"
	for currVertex := vertices[targetPos]; currVertex != nil && currVertex.prev != nil; currVertex = currVertex.prev {
		prevPos := currVertex.prev.p
		deltaX := currVertex.p.x - prevPos.x
		deltaY := currVertex.p.y - prevPos.y
		if deltaX == 1 {
			path[prevPos] = ">"
		} else if deltaX == -1 {
			path[prevPos] = "<"
		} else if deltaY == 1 {
			path[prevPos] = "v"
		} else if deltaY == -1 {
			path[prevPos] = "^"
		} else {
			panic("error")
		}
	}

	for y := 0; y < mapHeight; y++ {
		for x := 0; x < mapWidth; x++ {
			p := Pos{x: x, y: y}
			if s, ok := path[p]; ok {
				fmt.Print(s)
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}

	fmt.Printf("startPos: %v\t targetPos:\t%v\n", dist[startPos], dist[targetPos])
	fmt.Printf("path length:\t%v", len(path))
	fmt.Printf("targetPos:\t%v", targetPos)
}

func abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}

func findNeighbors(heights map[Pos]int, p Pos) Neighbors {
	selfHeight := heights[p]
	neigh := make(Neighbors, 0, 9)

	for y := p.y - 1; y <= p.y+1; y++ {
		for x := p.x - 1; x <= p.x+1; x++ {
			if y < 0 || y > mapHeight || x < 0 || x > mapWidth {
				continue
			}

			if abs(p.x-x)+abs(p.y-y) > 1 {
				continue
			}

			n := Pos{x, y}
			if n == p {
				continue
			}

			if heights[n]-selfHeight <= 1 {
				neigh = append(neigh, n)
			}
		}
	}

	return neigh
}
