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

	// var c byte
	// for c = 'a'; c <= 'z'; c++ {
	// 	fmt.Printf("letter %v with prio %v\n", string(c), priority(c))
	// }
	// for c = 'A'; c <= 'Z'; c++ {
	// 	fmt.Printf("letter %v with prio %v\n", string(c), priority(c))
	// }

	log.Println("AOC - 2022.12.03+2")

	inputBytes, err := os.ReadFile("../input")
	if err != nil {
		log.Fatalf("couldn't open input file. error %v", err)
		return
	}

	totalScore := 0

	input := strings.Split(string(inputBytes), "\n")
	for i := 0; i < len(input); i += 3 {
		elf1 := input[i]
		elf2 := input[i+1]
		elf3 := input[i+2]

		m := make(map[byte]interface{}, 100)
		for i := range elf1 {
			m[elf1[i]] = nil
		}
		m2 := make(map[byte]interface{}, 100)
		for i := range elf2 {
			c := elf2[i]
			if _, ok := m[c]; ok {
				m2[c] = nil
			}
		}
		for i := range elf3 {
			c := elf3[i]
			if _, ok := m2[c]; ok {
				totalScore += priority(c)
				break
			}
		}

	}

	fmt.Println("total score:", totalScore)
}
