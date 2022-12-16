package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Packet struct {
	data   []any
	parent *Packet
}

func (p Packet) String() string {
	return fmt.Sprint(p.data)
}

func main() {

	log.Println("AOC - 2022.12.12")

	inputBytes, err := os.ReadFile("../input")
	// inputBytes, err := os.ReadFile("../test_input")
	if err != nil {
		log.Fatalf("couldn't open input file. error %v", err)
		return
	}

	input := strings.Split(string(inputBytes), "\n")
	packets := make([]Packet, 0, len(input)/2)

	for _, line := range input {
		if line == "" {
			continue
		}

		rootPacket := Packet{
			data:   make([]any, 0),
			parent: nil,
		}
		currPacket := &rootPacket

		val := 0
		parsingDigit := false
		for idx := 1; idx < len(line); idx++ {
			if line[idx] == '[' {
				newPacket := Packet{
					data:   make([]any, 0, len(line)),
					parent: currPacket,
				}
				currPacket = &newPacket
			} else if line[idx] == ']' {
				if parsingDigit {
					currPacket.data = append(currPacket.data, val)
				}
				if idx != len(line)-1 {
					currPacket.parent.data = append(currPacket.parent.data, currPacket.data)
					parsingDigit = false
					val = 0
					currPacket = currPacket.parent
				}
			} else if line[idx] != ',' {
				parsingDigit = true
				digit := int(line[idx] - '0')
				val = 10*val + digit
			} else if line[idx] == ',' {
				if parsingDigit {
					currPacket.data = append(currPacket.data, val)
					parsingDigit = false
					val = 0
				}
			}
		}

		packets = append(packets, rootPacket)
	}

	sum := 0

	for idx := 0; idx < len(packets); idx += 2 {
		lPacket := packets[idx]
		rPacket := packets[idx+1]

		pairIdx := idx/2 + 1

		switch res := lPacket.compareTo(&rPacket); res {
		case Undecided:
			panic("all pairs should be either in or out of order")
		case InOrder:
			sum += pairIdx
		}
	}

	fmt.Printf("sum:\t%v", sum)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

type Result int

const (
	OutOfOrder = iota
	InOrder
	Undecided
)

func (l *Packet) compareTo(r *Packet) Result {

	var res Result = Undecided

	if len(l.data) == 0 && len(r.data) > 0 {
		// left side ran out of items - in order
		return InOrder
	} else if len(l.data) > 0 && len(r.data) == 0 {
		// right side ran out of items - out of order
		return OutOfOrder
	} else if len(l.data) == 0 && len(r.data) == 0 {
		return Undecided
		// panic("l.data and r.data are empty")
	}

	for i := 0; i < len(l.data); i++ {
		lElem := l.data[i]
		if i >= len(r.data) {
			// right side ran out of items - out of order
			return OutOfOrder
		}
		rElem := r.data[i]

		switch l := lElem.(type) {
		case []any:
			switch r := rElem.(type) {
			case []any:
				lPacket := packetFromData(l)
				rPacket := packetFromData(r)
				res = lPacket.compareTo(&rPacket)

				if res != Undecided {
					return res
				}
			case int:
				lPacket := packetFromData(l)
				rPacket := packetFromData([]any{r})
				res = lPacket.compareTo(&rPacket)

				if res != Undecided {
					return res
				}
			}
		case int:
			switch r := rElem.(type) {
			case int:
				if l < r {
					return InOrder
				} else if l > r {
					return OutOfOrder
				}
			case []any:
				lPacket := packetFromData([]any{l})
				rPacket := packetFromData(r)
				res = lPacket.compareTo(&rPacket)

				if res != Undecided {
					return res
				}
			}
		}
	}

	return InOrder
}

func packetFromData(data []any) Packet {
	p := Packet{
		data: data,
	}

	return p
}
