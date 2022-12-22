package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	next         *Node
	prev         *Node
	originalNext *Node
	val          int
}

func main() {

	log.Println("AOC - 2022.12.20")

	inputBytes, err := os.ReadFile("../input")
	if err != nil {
		log.Fatalf("couldn't open input file. error %v", err)
		return
	}

	arr := make([]*Node, 0)
	firstNode := &Node{}

	prevNode := firstNode
	input := strings.Split(string(inputBytes), "\n")
	for _, line := range input {
		n, _ := strconv.Atoi(line)
		node := Node{
			prev: prevNode,
			val:  811589153 * n,
		}
		prevNode.next = &node
		prevNode.originalNext = &node
		prevNode = &node

		arr = append(arr, &node)
	}

	firstNode = firstNode.next

	prevNode.next = firstNode
	firstNode.prev = prevNode

	currNode := arr[0]
	for i := 0; i < 2*len(arr); i++ {
		fmt.Printf("%d, ", currNode.val)
		currNode = currNode.next
	}
	fmt.Println()

	for i := 0; i < 10; i++ {

		for _, n := range arr {
			n.move(len(arr))
			// currNode = arr[0]
			// for i := 0; i < len(arr); i++ {
			// 	fmt.Printf("%d, ", currNode.val)
			// 	currNode = currNode.next
			// }
			// fmt.Println()
		}
	}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d, ", currNode.val)
		currNode = currNode.next
	}
	fmt.Println()

	currNode = firstNode
	for currNode.val != 0 {
		currNode = currNode.next
	}

	endList := make([]int, 0, len(arr))
	endList = append(endList, currNode.val)
	currNode = currNode.next

	for currNode.val != 0 {
		endList = append(endList, currNode.val)
		currNode = currNode.next
	}

	n := len(endList)
	r1, r2, r3 := endList[1000%n], endList[2000%n], endList[3000%n]
	fmt.Println(r1 + r2 + r3)

}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	} else {
		return n
	}
}

func (n *Node) move(listLength int) {
	n.prev.next = n.next
	n.next.prev = n.prev

	for i := 0; i < abs(n.val)%(listLength-1); i++ {
		if n.val > 0 {
			n.prev = n.next
			n.next = n.next.next
		} else {
			n.next = n.prev
			n.prev = n.prev.prev
		}
	}

	n.prev.next = n
	n.next.prev = n
}
