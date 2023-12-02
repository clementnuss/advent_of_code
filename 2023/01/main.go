package main

import (
	"log"
	"os"
	"regexp"
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

	res := 0
	for _, line := range strings.Split(string(inputBytes), "\n") {
		firstDigit, lastDigit := -1, -1
		for _, char := range line {
			if char >= '0' && char <= '9' {
				if firstDigit == -1 {
					firstDigit = int(char - '0')
				} else {
					lastDigit = int(char - '0')
				}
			}
		}
		if lastDigit == -1 {
			lastDigit = firstDigit
		}
		res += 10*firstDigit + lastDigit
	}

	log.Printf("part 1: %v", res)

	strToInt := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	firstDigitRgx, _ := regexp.Compile(`(one|two|three|four|five|six|seven|eight|nine|\d){1}?`)
	lastDigitRgx, _ := regexp.Compile(`(?:.*)(one|two|three|four|five|six|seven|eight|nine|\d){1}`)

	res = 0
	for _, line := range strings.Split(string(inputBytes), "\n") {
		var firstDigit, lastDigit int
		match := firstDigitRgx.FindString(line)
		if val, ok := strToInt[match]; ok {
			firstDigit = val
		} else {
			firstDigit, _ = strconv.Atoi(match)
		}

		match = lastDigitRgx.FindStringSubmatch(line)[1]
		if val, ok := strToInt[match]; ok {
			lastDigit = val
		} else {
			lastDigit, _ = strconv.Atoi(match)
		}

		res += 10*firstDigit + lastDigit
	}

	log.Printf("part 2: %v", res)
}
