package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func priority(c byte) int {
	if c >= 'a' {

		return int(c) - int('a') + 1
	} else {
		return 26 + int(c) - int('A') + 1
	}
}

func main() {

	var c byte
	for c = 'a'; c <= 'z'; c++ {
		fmt.Printf("letter %v with prio %v\n", string(c), priority(c))
	}
	for c = 'A'; c <= 'Z'; c++ {
		fmt.Printf("letter %v with prio %v\n", string(c), priority(c))
	}

	log.Println("AOC - 2022.12.03+1")

	inputBytes, err := os.ReadFile("../input")
	if err != nil {
		log.Fatalf("couldn't open input file. error %v", err)
		return
	}

	totalScore := 0

	input := strings.Split(string(inputBytes), "\n")
	for _, rucksack := range input {
		half := len(rucksack) / 2
		// scan first half
		m := make(map[byte]interface{}, half)
		for i := 0; i < half; i++ {
			m[rucksack[i]] = nil
		}
		// scan second hald
		for i := half; i < len(rucksack); i++ {
			if _, ok := m[rucksack[i]]; ok {
				totalScore += priority(rucksack[i])
				break
			}

		}
	}

	fmt.Println("total score:", totalScore)
}
