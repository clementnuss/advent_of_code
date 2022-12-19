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

	totalScore := 0
	minX, minY := 0, 0
	maxX, maxY := 0, 0

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

		if newMinX := sensor.x - manhDist; newMinX < minX {
			minX = newMinX
		}

		if newMinY := sensor.y - manhDist; newMinY < minY {
			minY = newMinY
		}

		if newMaxX := sensor.x + manhDist; newMaxX > maxX {
			maxX = newMaxX
		}

		if newMaxY := sensor.y + manhDist; newMaxY > maxY {
			maxY = newMaxY
		}
	}

	y := 2000000

	for x := minX; x < maxX; x++ {
		p := Pos{x, y}
		for sensor, dist := range sensors {
			if p.manhattanDistance(&sensor) <= dist {
				if _, ok := beacons[p]; !ok {
					totalScore++
					break
				}
			}
		}
	}

	fmt.Println("total score:", totalScore)
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
