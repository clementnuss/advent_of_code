package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	log.Println("AOC - 2023.12.")

	inputBytes, err := os.ReadFile("../input")
	if err != nil {
		log.Fatalf("couldn't open input file. error %v", err)
		return
	}

	totalScore := 0

	input := strings.Split(string(inputBytes), "\n")
	for _, line := range input {

	}

	fmt.Println("total score:", totalScore)
}
