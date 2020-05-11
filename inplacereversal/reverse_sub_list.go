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
	start, end := pos[0], pos[1]
	if start == end {
		return head
	}
	current := head
	var prev *MyNode1
	// We are taking three parts of the linked list in consideration.
	// the first part is before the index 'start'
	// the second part is the sub-list between 'start' and 'end'
	// the third part is after index 'end'
	// this is the first part of the linked list, 'current' will points to 'start'th node
	for current != nil && current.value != start {
		prev = current
		current = current.next
	}
	// last node of the first part of linked list
	lastNodeOfFirstPart := prev
	// we will reverse the second part, so the first node we meet in the second part, it becomes the last node of sub list
	lastNodeOfSubList := current
	// reverse the second part of the linked list
	for current != nil && prev.value != end { //use prev as base check, because it's the last node of the sublist should be inclusive
		next := current.next
		current.next = prev
		prev = current
		current = next // move
	}
	// we connect with the first part
	if lastNodeOfFirstPart != nil {
		lastNodeOfFirstPart.next = prev
	} else {
		head = prev
	}
	// we connect with the last part
	lastNodeOfSubList.next = current
	return head

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
	head.next = &MyNode1{value: 3}
	head.next.next = &MyNode1{value: 9}
	head.next.next.next = &MyNode1{value: 4}
	head.next.next.next.next = &MyNode1{value: 5}
	head = reverseSubList(head, []int{3, 4})
	printMyNode1(head)

}
