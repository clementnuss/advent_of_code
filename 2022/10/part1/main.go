package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instr func(*CPU) bool

type CPU struct {
	X            int
	currentInstr Instr
	instrChan    chan Instr
}

func noop(*CPU) bool {
	return true
}

func addx(x int) Instr {
	i := 0
	return func(c *CPU) bool {
		i++
		if i == 2 {
			c.X += x
			return true
		}
		return false
	}
}

func main() {

	log.Println("AOC - 2022.12.10+1")

	inputBytes, err := os.ReadFile("../input")
	if err != nil {
		log.Fatalf("couldn't open input file. error %v", err)
		return
	}

	c := &CPU{
		X:            1,
		currentInstr: nil,
		instrChan:    make(chan Instr, 1e3),
	}

	totalScore := 0

	input := strings.Split(string(inputBytes), "\n")
	for _, instr := range input {
		if instr == "noop" {
			c.instrChan <- noop
			continue
		}

		spl := strings.Split(instr, " ")
		if spl[0] == "addx" {
			x, _ := strconv.Atoi(spl[1])
			c.instrChan <- addx(x)
		} else {
			panic("shouldn't be here")
		}
	}

	c.currentInstr = <-c.instrChan

	for cycle := 1; ; cycle++ {
		if (cycle-20)%40 == 0 {
			fmt.Printf("cycle:\t%v\tx:%v\n", cycle, c.X)
			totalScore += cycle * c.X
		}

		if c.currentInstr(c) {
			if len(c.instrChan) == 0 {
				break
			}

			c.currentInstr = <-c.instrChan
		}
	}

	fmt.Println("total score:", totalScore)
}
