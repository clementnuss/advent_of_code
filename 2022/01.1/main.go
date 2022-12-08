package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	log.Println("AOC - 2022.12.01+1")

	inputBytes, err := os.ReadFile("input")
	if err != nil {
		log.Fatalf("couldn't open input file. error %v", err)
		return
	}

	currentTotal, max := 0, -1
	inputStr := string(inputBytes)
	for _, cal := range strings.Split(inputStr, "\n") {
		if cal == "" {
			if currentTotal > max {
				max = currentTotal
			}

			currentTotal = 0
			continue
		}

		c, err := strconv.Atoi(cal)
		if err != nil {
			log.Fatalf("wrong input: %s", cal)
			return
		}

		currentTotal += c
	}

	log.Printf("max cal: %v", max)

}
