package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Color int64

const (
	Red Color = iota
	Green
	Blue
)

type Draw struct {
	red   int
	green int
	blue  int
}

type Game struct {
	id    int
	draws []Draw
}

func main() {

	log.Println("AOC - 2023.12.02")

	inputBytes, err := os.ReadFile("./input")
	if err != nil {
		log.Fatalf("couldn't open input file. error %v", err)
		return
	}

	games := make([]Game, 100)
	input := strings.Split(string(inputBytes), "\n")
	for i, line := range input {
		column := strings.Index(line, ":")
		id, _ := strconv.Atoi(line[len("Game "):column])
		games[i].id = id
		line = line[column+len(": "):]
		draws := strings.Split(line, "; ")
		for _, draw := range draws {
			games[i].draws = append(games[i].draws, *parseDraw(draw))
		}
	}

	res1 := 0
	maxDraw := Draw{
		red:   12,
		green: 13,
		blue:  14,
	}

	for _, game := range games {
		valid := true
		for _, dr := range game.draws {
			if dr.red > maxDraw.red || dr.green > maxDraw.green || dr.blue > maxDraw.blue {
				valid = false
			}
		}
		if valid {
			res1 += game.id
		}
	}
	fmt.Println("part1:", res1)

	res2 := 0
	for _, game := range games {
		minDraw := Draw{}
		for _, dr := range game.draws {
			minDraw.red = max(minDraw.red, dr.red)
			minDraw.green = max(minDraw.green, dr.green)
			minDraw.blue = max(minDraw.blue, dr.blue)
		}
		res2 += minDraw.red * minDraw.green * minDraw.blue
	}
	fmt.Println(res2)
}

func parseDraw(s string) (dr *Draw) {
	dr = new(Draw)
	spl := strings.Split(s, ", ")

	for len(spl) > 0 {
		flds := strings.Fields(spl[0])
		count, _ := strconv.Atoi(flds[0])
		col := flds[1]
		switch col {
		case "red":
			dr.red = count
		case "green":
			dr.green = count
		case "blue":
			dr.blue = count
		}
		spl = spl[1:]
	}
	return
}
