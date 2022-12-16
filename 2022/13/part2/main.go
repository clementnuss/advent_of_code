package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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

	log.Println("AOC - 2022.12.13")

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

	res := 1
	div2 := Packet{
		data: []any{[]any{2}},
	}
	div6 := Packet{
		data: []any{[]any{6}},
	}

	packets = append(packets, div2, div6)

	sort.Slice(packets, func(i, j int) bool {
		lPacket := packets[i].data
		rPacket := packets[j].data
		switch res := betterCompare(lPacket, rPacket); res {
		case InOrder:
			return true
		case OutOfOrder:
			return false
		default:
			return false
		}
	})

	for idx := 0; idx < len(packets); idx++ {
		fmt.Printf("packet %v\n", packets[idx])
		if betterCompare(packets[idx].data, div2.data) == Undecided ||
			betterCompare(packets[idx].data, div6.data) == Undecided {
			res *= idx + 1
			fmt.Printf("packet %v\n", idx+1)
		}
	}

	fmt.Printf("sum:\t%v", res)
}

type Result int

const (
	OutOfOrder = iota
	InOrder
	Undecided
)

// func (src *Packet) Clone() Packet {
// 	return Packet{
// 		data: recursiveClone(src.data),
// 	}
// }

// func recursiveClone(srcData []any) []any {
// 	retData := make([]any, 0)

// 	if len(srcData) == 0 {
// 		return make([]any, 0)
// 	}

// 	for _, val := range srcData {
// 		switch v := val.(type) {
// 		case int:
// 			retData = append(retData, v)
// 		case []any:
// 			retData = append(retData, recursiveClone(v))
// 		}
// 	}
// 	return retData
// }

func betterCompare(l, r []any) Result {
	i := 0
	for ; i < len(l) && i < len(r); i++ {
		var res Result = Undecided

		lElem := l[0]
		rElem := r[0]

		switch lVal := lElem.(type) {
		case int:
			switch rVal := rElem.(type) {
			case int:
				if lVal == rVal {
					continue
				} else if lVal < rVal {
					return InOrder
				} else if lVal > rVal {
					return OutOfOrder
				}
			case []any:
				lArr := []any{lVal}
				res = betterCompare(lArr, rVal)
			}
		case []any:
			switch rVal := rElem.(type) {
			case int:
				rArr := []any{rVal}
				res = betterCompare(lVal, rArr)
			case []any:
				res = betterCompare(lVal, rVal)
			}
		}
		if res != Undecided {
			return res
		}
	}

	if i == len(l) && i == len(r) {
		return Undecided
	} else if i == len(l) { // l finished first : in order
		return InOrder
	} else {
		return OutOfOrder
	}
}

// func (l *Packet) compareTo(r *Packet) Result {

// 	var res Result = Undecided

// 	if len(l.data) == 0 && len(r.data) > 0 {
// 		// left side ran out of items - in order
// 		return InOrder
// 	} else if len(l.data) > 0 && len(r.data) == 0 {
// 		// right side ran out of items - out of order
// 		return OutOfOrder
// 	} else if len(l.data) == 0 && len(r.data) == 0 {
// 		return Undecided
// 	}

// 	for len(l.data) > 0 {
// 		if len(r.data) == 0 {
// 			// right side ran out of items - out of order
// 			return OutOfOrder
// 		}

// 		lElem := l.data[0]
// 		rElem := r.data[0]

// 		switch lVal := lElem.(type) {
// 		case []any:
// 			switch rVal := rElem.(type) {
// 			case []any:
// 				lPacket := Packet{data: lVal}
// 				rPacket := Packet{data: rVal}
// 				res = lPacket.compareTo(&rPacket)

// 				if res != Undecided {
// 					return res
// 				}

// 				l.data = l.data[1:]
// 				r.data = r.data[1:]

// 			case int:
// 				lPacket := l
// 				r.data = r.data[1:]
// 				rPacket := Packet{
// 					data: []any{[]any{rVal}, r.data},
// 				}
// 				res = lPacket.compareTo(&rPacket)

// 				if res != Undecided {
// 					return res
// 				}
// 			}
// 		case int:
// 			switch rVal := rElem.(type) {
// 			case int:
// 				l.data = l.data[1:]
// 				r.data = r.data[1:]

// 				if lVal < rVal {
// 					return InOrder
// 				} else if lVal > rVal {
// 					return OutOfOrder
// 				}
// 			case []any:
// 				l.data = l.data[1:]
// 				lPacket := Packet{
// 					data: []any{[]any{lVal}, l.data},
// 				}
// 				rPacket := r
// 				res = lPacket.compareTo(rPacket)

// 				if res != Undecided {
// 					return res
// 				}
// 			}
// 		}
// 	}

// 	return Undecided
// }
