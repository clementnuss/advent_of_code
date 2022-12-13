package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items   []int
	op      func(int) int
	test    int
	ifTrue  int
	ifFalse int
}

func (m *Monkey) throwTo(item int) int {
	if item%m.test == 0 {
		return m.ifTrue
	} else {
		return m.ifFalse
	}
}

func (m *Monkey) append(item int) {
	m.items = append(m.items, item)
}

func main() {

	log.Println("AOC - 2022.12.04+1")

	inputBytes, err := os.ReadFile("../input")
	// inputBytes, err := os.ReadFile("./test_input")
	if err != nil {
		log.Fatalf("couldn't open input file. error %v", err)
		return
	}

	monkeys := make([]*Monkey, 0, 10)

	var currMonkey *Monkey

	input := strings.Split(string(inputBytes), "\n")
	for _, line := range input {
		if line == "" {
			continue
		}

		if strings.Contains(line, "Monkey") {
			currMonkey = &Monkey{
				items: make([]int, 0, 20),
			}
			monkeys = append(monkeys, currMonkey)
		}

		if list := strings.TrimPrefix(line, "  Starting items:"); list != line {
			itemsList := strings.Split(list, ",")
			for _, item := range itemsList {
				itemLevel, _ := strconv.Atoi(strings.TrimSpace(item))
				currMonkey.items = append(currMonkey.items, itemLevel)
			}
		}

		if op := strings.TrimPrefix(line, "  Operation: new = "); op != line {
			fields := strings.Fields(op)
			switch fields[1] {
			case "+":
				currMonkey.op = func(i int) int {
					if fields[2] == "old" {
						return i + i
					}

					incr, _ := strconv.Atoi(fields[2])

					return i + incr
				}
			case "*":
				currMonkey.op = func(i int) int {
					if fields[2] == "old" {
						return i * i
					}

					incr, _ := strconv.Atoi(fields[2])

					return i * incr
				}
			}
			continue
		}

		if test := strings.TrimPrefix(line, "  Test: divisible by "); test != line {
			t, _ := strconv.Atoi(test)
			currMonkey.test = t
		}

		if ifTrue := strings.TrimPrefix(line, "    If true: throw to monkey "); ifTrue != line {
			monkey, _ := strconv.Atoi(ifTrue)
			currMonkey.ifTrue = monkey
		}

		if ifFalse := strings.TrimPrefix(line, "    If false: throw to monkey "); ifFalse != line {
			monkey, _ := strconv.Atoi(ifFalse)
			currMonkey.ifFalse = monkey
		}
	}

	monkeyBusiness := make([]int, len(monkeys))

	for round := 1; round < 21; round++ {
		for idx, currMonkey := range monkeys {
			for _, item := range currMonkey.items {
				monkeyBusiness[idx]++

				newLevel := currMonkey.op(item) / 3
				throwToMonkey := currMonkey.throwTo(newLevel)
				monkeys[throwToMonkey].append(newLevel)
			}

			currMonkey.items = currMonkey.items[:0]
		}
	}

	sort.IntSlice(monkeyBusiness).Sort()

	n := len(monkeys)

	fmt.Println("business:", monkeyBusiness[n-1]*monkeyBusiness[n-2])
}

func (m *Monkey) String() string {
	return fmt.Sprintf("monkey with items %v\n", m.items)

}
