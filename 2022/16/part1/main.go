package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	log.Println("AOC - 2022.12.16")

	inputBytes, err := os.ReadFile("../input")
	// inputBytes, err := os.ReadFile("../test_input")
	if err != nil {
		log.Fatalf("couldn't open input file. error %v", err)
		return
	}

	r := regexp.MustCompile(`^Valve\s(\w{2}).*=(\d+);.*valves?\s(.*)`)

	flow := make(map[string]int)
	tunnels := make(map[string]map[string]int)

	input := strings.Split(string(inputBytes), "\n")
	for _, line := range input {
		submatch := r.FindStringSubmatch(line)
		v := submatch[1]
		vFlow, _ := strconv.Atoi(submatch[2])
		edges := strings.Split(submatch[3], ", ")

		if vFlow > 0 {
			flow[v] = vFlow
		}

		tunnels[v] = make(map[string]int, 0)
		for _, vv := range edges {
			tunnels[v][vv] = 1
		}
	}

	bitMask := make(map[string]int)

	{
		i := 1
		for v := range flow {
			bitMask[v] = i
			i <<= 1
		}
	}

	n := len(tunnels)
	dist := make(map[string]map[string]int, n)

	for v := range tunnels {
		dist[v] = make(map[string]int, n)

		for w := range tunnels {
			if d, ok := tunnels[v][w]; ok {
				dist[v][w] = d
			} else {
				dist[v][w] = 1e10
			}
		}
	}

	// floyd warshall
	for k := range tunnels {
		for i := range tunnels {
			for j := range tunnels {
				newDist := dist[i][k] + dist[k][j]
				if newDist < dist[i][j] {
					dist[i][j] = newDist
				}
			}
		}
	}

	maxForState := make(map[int]int)
	bfs("AA", 30, 0, 0, maxForState, dist, flow, bitMask)

	max := 0
	for _, m := range maxForState {
		if m > max {
			max = m
		}
	}

	fmt.Println("max pressure relief:", max)
}

func bfs(v string, timeLeft, pressureRelief, state int, maxForState map[int]int,
	dist map[string]map[string]int, flow map[string]int, bitMask map[string]int) {
	for w, d := range dist[v] {
		if (bitMask[w]&state > 0) || flow[w] == 0 {
			continue
		}

		newTime := timeLeft - d - 1
		if newTime < 0 {
			continue
		}

		newPressureRelief := pressureRelief + newTime*flow[w]
		newState := state | bitMask[w]

		if oldVal, ok := maxForState[newState]; ok {
			if newPressureRelief > oldVal {
				maxForState[newState] = newPressureRelief
			}
		} else {
			maxForState[newState] = newPressureRelief
		}

		bfs(w, newTime, newPressureRelief, newState, maxForState, dist, flow, bitMask)
	}
}
