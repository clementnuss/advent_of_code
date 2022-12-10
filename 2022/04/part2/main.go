package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Assignment struct {
	start int
	end   int
}

func (a *Assignment) size() int {
	return a.end - a.start + 1
}

func (a *Assignment) contained(o *Assignment) bool {
	if a.size() < o.size() {
		o, a = a, o // we swap the pointers and consider a to be the largest
	}

	if a.start <= o.start && a.end >= o.end {
		return true
	}
	return false
}

func (a *Assignment) overlap(b *Assignment) bool {

	return !(b.end < a.start || b.start > a.end)

}

func (a *Assignment) parse(s string) {
	spl := strings.Split(s, "-")
	a.start, _ = strconv.Atoi(spl[0])
	a.end, _ = strconv.Atoi(spl[1])
}

func main() {

	log.Println("AOC - 2022.12.04+1")

	inputBytes, err := os.ReadFile("../input")
	if err != nil {
		log.Fatalf("couldn't open input file. error %v", err)
		return
	}

	totalScore := 0

	input := strings.Split(string(inputBytes), "\n")
	for _, line := range input {
		spl := strings.Split(line, ",")
		var a1, a2 Assignment
		a1.parse(spl[0])
		a2.parse(spl[1])
		if a1.overlap(&a2) {
			totalScore++
		}
	}

	fmt.Println("total score:", totalScore)
}
