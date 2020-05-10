package main

import (
	"fmt"
)

// Node simple node
type Node struct {
	value int
	next  *Node
}

/*
	Reverse a singly linkedlist in an inpace-way
*/
func reverse(head *Node) *Node {
	current := head
	var prev *Node
	for current != nil {
		next := current.next //temporarily save the next node
		current.next = prev
		prev = current
		current = next // move pointer
	}
	return prev

}

func printList(head *Node) {
	current := head
	for current != nil {
		fmt.Print(current.value)
		fmt.Print(" ")
		current = current.next
	}
	fmt.Println()
}

func main() {
	head := &Node{value: 2}
	head.next = &Node{value: 4}
	head.next.next = &Node{value: 6}
	head.next.next.next = &Node{value: 8}
	head.next.next.next.next = &Node{value: 10}
	fmt.Println("Before")
	printList(head)
	head = reverse(head)
	fmt.Println("After")
	printList(head)
}
