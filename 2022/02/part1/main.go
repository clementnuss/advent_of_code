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

func (self *Item) pointsAgainst(opponent *Item) int {
	if *self == *opponent {
		return 3
	}

	fmt.Println(*self - *opponent)
	fmt.Println(*self - *opponent%3)
	if *self-*opponent%3 == 1 {
		return 6
	}

	return 0

	// switch *i {
	// case Rock:
	// 	if *o == Scissors {
	// 		return 6
	// 	}
	// case Paper:
	// 	if *o == Rock {
	// 		return 6
	// 	}
	// case Scissors:
	// 	if *o == Paper {
	// 		return 6
	// 	}
	// }
	// return 0
}

var m = map[byte]Item{
	'A': Rock,
	'B': Paper,
	'C': Scissors,
	'X': Rock,
	'Y': Paper,
	'Z': Scissors,
}

func main() {

	log.Println("AOC - 2022.12.02+1")

	inputBytes, err := os.ReadFile("../input")
	if err != nil {
		log.Fatalf("couldn't open input file. error %v", err)
		return
	}

	totalScore := 0

	input := strings.Split(string(inputBytes), "\n")
	for _, round := range input {
		opponent, self := m[round[0]], m[round[2]]
		totalScore += int(self)
		totalScore += self.pointsAgainst(&opponent)
	}

	fmt.Println("total score:", totalScore)
}
