package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	log.Println("AOC - 2022.12.04+1")

	inputBytes, err := os.ReadFile("../input")
	if err != nil {
		log.Fatalf("couldn't open input file. error %v", err)
		return
	}

	counts := make(map[byte]int, 26)

	for idx, char := range inputBytes {
		if idx >= 14 {
			if len(counts) == 14 {
				fmt.Println("current index: ", idx)
				return
			}

			oldChar := inputBytes[idx-14]
			counts[oldChar]--
			if counts[oldChar] == 0 {
				delete(counts, oldChar)
			}
		}
		counts[char]++
	}

}
