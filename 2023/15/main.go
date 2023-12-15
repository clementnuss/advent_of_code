package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	orderedmap "github.com/wk8/go-ordered-map/v2"
)

type Hasher struct {
	currentVal uint
}

const debug bool = false

func (h *Hasher) update(l string) uint {
	for _, c := range l {
		h.currentVal += uint(c)
		h.currentVal *= 17
		h.currentVal &= 0xff
	}
	return h.currentVal
}

func (h *Hasher) reset() {
	h.currentVal = 0
}

func main() {
	log.Println("AOC - 2023.12.15")

	inputBytes, err := os.ReadFile("input")
	// inputBytes, err := os.ReadFile("example_input")
	if err != nil {
		log.Fatalf("couldn't open input file. error %v", err)
		return
	}

	res1 := uint(0)
	h := new(Hasher)
	re := regexp.MustCompile(`([a-z]+)(-|=\d?)`)

	boxes := make([]*orderedmap.OrderedMap[string, int], 256)
	for _, line := range strings.Split(string(inputBytes), ",") {
		subMatch := re.FindStringSubmatch(line)
		l := subMatch[1]
		lh := h.update(l)
		h.reset()
		res1 += lh

		if boxes[lh] == nil {
			boxes[lh] = orderedmap.New[string, int]()
		}

		b := boxes[lh]

		switch subMatch[2][:1] {
		case "-":
			b.Delete(l)

		case "=":
			lens := int(subMatch[2][1:][0] - '0')
			if _, ok := b.Get(l); ok {
				p := b.GetPair(l)
				p.Value = (lens)
				// om.Set(label, lens)
			} else {
				b.Set(l, lens)
				// _, _ = om.GetAndMoveToBack(label)
			}
		}
	}
	fmt.Println(res1)

	res2 := 0
	for i, om := range boxes {
		if om == nil || om.Len() == 0 {
			continue
		}

		if debug {
			fmt.Printf("\nlabel: %2v box: ", i)
		}
		slot := 1
		for p := om.Oldest(); p != nil; p = p.Next() {
			if debug {
				fmt.Printf("%6s %v ", p.Key, p.Value)
			}
			res2 += (i + 1) * (slot) * (p.Value)
			slot++
		}
	}
	fmt.Println()

	fmt.Println(res2)
}
