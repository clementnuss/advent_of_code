package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	log.Println("AOC - 2022.12.01+2")

	inputBytes, err := os.ReadFile("input")
	if err != nil {
		log.Fatalf("couldn't open input file. error %v", err)
		return
	}

	var totals []int

	currTotal := 0
	inputStr := string(inputBytes)
	for _, cal := range strings.Split(inputStr, "\n") {
		if cal == "" {
			totals = append(totals, currTotal)
			currTotal = 0
			continue
		}

		c, err := strconv.Atoi(cal)
		if err != nil {
			log.Fatalf("wrong input: %s", cal)
			return
		}

		currTotal += c
	}

	sortedTotal := sort.IntSlice(totals)
	sortedTotal.Sort()

	max3Total := 0
	for _, cal := range sortedTotal[len(sortedTotal)-3:] {
		max3Total += cal
	}

	log.Printf("sum of max 3 cals: %v", max3Total)

}
