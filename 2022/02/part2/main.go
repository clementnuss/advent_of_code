package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Item int

const (
	Rock Item = iota + 1
	Paper
	Scissors
)

var m = map[byte]Item{
	'A': Rock,
	'B': Paper,
	'C': Scissors,
}

func (self *Item) pointsAgainst(opponent *Item) int {
	if *self == *opponent {
		return 3
	}

	// fmt.Println(*self - *opponent)
	// fmt.Println(*self - *opponent%3)
	if *self-*opponent%3 == 1 {
		return 6
	}

	return 0
}

func main() {

	log.Println("AOC - 2022.12.02+2")

	inputBytes, err := os.ReadFile("../input")
	if err != nil {
		log.Fatalf("couldn't open input file. error %v", err)
		return
	}

	totalScore := 0

	input := strings.Split(string(inputBytes), "\n")
	for _, round := range input {
		opponent := m[round[0]]
		var self Item
		switch round[2] {
		case 'X': // need to lose
			self = Item((opponent+2-1)%3 + 1)
		case 'Y':
			self = opponent
		case 'Z':
			self = Item((opponent)%3 + 1)
		}

		totalScore += int(self)
		totalScore += self.pointsAgainst(&opponent)
	}

	fmt.Println("total score:", totalScore)
}
