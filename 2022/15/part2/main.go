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

func main() {

	log.Println("AOC - 2022.12.15")

	// inputBytes, err := os.ReadFile("../test_input")
	inputBytes, err := os.ReadFile("../input")
	if err != nil {
		log.Fatalf("couldn't open input file. error %v", err)
		return
	}

	result := 0
	sensors := make(map[Pos]int)
	beacons := make(map[Pos]any)

	input := strings.Split(string(inputBytes), "\n")
	for _, line := range input {
		sensorIdx := strings.Index(line, "x")
		columnIdx := strings.Index(line, ":")
		sensor := parsePos(line[sensorIdx:columnIdx])

		beaconIdx := strings.LastIndex(line, "x")
		beacon := parsePos(line[beaconIdx:])

		manhDist := sensor.manhattanDistance(&beacon)
		sensors[sensor] = manhDist
		beacons[beacon] = nil
	}

	searchArea := 4000000

	for s1, d1 := range sensors {
		for x := s1.x - (d1 + 1); x < s1.x+(d1+1); x++ {
			delta := d1 + 1 - abs(s1.x-x)
			if x < 0 || x > searchArea {
				continue
			}

			for _, y := range []int{s1.y - delta, s1.y + delta} {
				if y < 0 || y > searchArea {
					continue
				}
				p := Pos{x, y}
				found := true
				for s2, d2 := range sensors {
					if s2.manhattanDistance(&p) <= d2 {
						found = false
						continue
					}
				}

				if found {
					fmt.Println("found the one! ", p)
					fmt.Println("frequency:", searchArea*p.x+p.y)
				}

			}
		}
	}

	fmt.Println("total score:", result)
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	} else {
		return x
	}
}

func (p *Pos) manhattanDistance(o *Pos) int {
	return abs(p.x-o.x) + abs(p.y-o.y)
}

func parsePos(s string) Pos {
	spl := strings.Split(strings.TrimSpace(s), ",")
	x := spl[0][2:]
	y := spl[1][3:]

	xVal, _ := strconv.Atoi(x)
	yVal, _ := strconv.Atoi(y)

	return Pos{xVal, yVal}
}
