package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Stack []byte

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(c byte) {
	*s = append(*s, c)
}

func (s *Stack) Pop() byte {
	if s.IsEmpty() {
		return ' '
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
	}
}

func parseStacks(input []string, index int) []Stack {

	stackNumber := (len(input[index-1]) + 2) / 4
	stacks := make([]Stack, stackNumber)

	for i := index - 2; i >= 0; i-- {
		for j := 1; j < len(input[i]); j += 4 {
			if input[i][j] != ' ' {
				stacks[j/4].Push(input[i][j])
			}
		}
	}
	return stacks
}

func main() {

	log.Println("AOC - 2022.12.05+1")

	inputBytes, err := os.ReadFile("../input")
	if err != nil {
		log.Fatalf("couldn't open input file. error %v", err)
		return
	}

	var stackList []Stack

	input := strings.Split(string(inputBytes), "\n")
	for idx, line := range input {
		if stackList == nil {
			if len(line) == 0 { // one line before the stack definitions
				stackList = parseStacks(input, idx)
			}
			continue
		}
		instr := strings.Split(line, " ")
		count, _ := strconv.Atoi(instr[1])
		from, _ := strconv.Atoi(instr[3])
		to, _ := strconv.Atoi(instr[5])

		for ; count > 0; count-- {
			stackList[to-1].Push(stackList[from-1].Pop())
		}

	}

	for _, stack := range stackList {
		fmt.Print(string(stack.Pop()))
	}

}
