package main

import (
	"fmt"
)

// MyNode1 singly linked node
type MyNode1 struct {
	value int
	next  *MyNode1
}

/*
Given the head of a LinkedList and two positions ‘p’ and ‘q’, reverse the LinkedList from position ‘p’ to ‘q’.
*/
func reverseSubList(head *MyNode1, pos []int) *MyNode1 {
	return nil
}
func printMyNode1(head *MyNode1) {
	current := head
	for current != nil {
		fmt.Print(current.value)
		fmt.Print(" ")
		current = current.next
	}
	fmt.Println()
}

func main() {
	head := &MyNode1{value: 1}
	head.next = &MyNode1{value: 2}
	head.next.next = &MyNode1{value: 3}
	head.next.next.next = &MyNode1{value: 4}
	head.next.next.next.next = &MyNode1{value: 5}
	head = reverseSubList(head, []int{2, 4})

}
