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
	i := 0
	// first part of the linked list, skipping start - 1 nodes, current will point to 'start'th node
	for current != nil && i < start-1 {
		prev = current
		current = current.next
		i++
	}
	// we are consider three parts of the linked list
	// the first part is before index 'start', the second part is between 'start' and 'end', the last part after index 'end'
	lastNodeOfFirstPart := prev
	lastNodeOfSubList := current
	i = 0
	// sub-list part of the linked list
	for current != nil && i < end-start+1 {
		next := current.next
		current.next = prev
		prev = current // prev is the last node of this part
		current = next
		i++
	}
	// connect with the first part
	if lastNodeOfFirstPart != nil {
		lastNodeOfFirstPart.next = prev
	} else {
		head = prev // start from the first node
	}
	// connect with the last part
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
	head.next = &MyNode1{value: 2}
	head.next.next = &MyNode1{value: 3}
	head.next.next.next = &MyNode1{value: 4}
	head.next.next.next.next = &MyNode1{value: 5}
	head = reverseSubList(head, []int{2, 4})
	printMyNode1(head)

}
