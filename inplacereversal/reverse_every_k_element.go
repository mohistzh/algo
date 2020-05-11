package main

import (
	"fmt"
)

// MyNode2 singly linked node
type MyNode2 struct {
	value int
	next  *MyNode2
}

/*
Given the head of a LinkedList and a number ‘k’, reverse every ‘k’ sized sub-list starting from the head.

If, in the end, you are left with a sub-list with less than ‘k’ elements, reverse it too.
*/
func reverseEveryKElements(head *MyNode2, k int) *MyNode2 {
	// base check
	if k <= 1 || head == nil {
		return head
	}
	current := head
	var prev *MyNode2
	for {
		lastNodeOfPrevPart := prev
		lastNodeOfSubList := current
		count := 0
		// reverse sublist
		for current != nil && count < k {
			next := current.next
			current.next = prev
			prev = current
			current = next
			count++
		}
		if lastNodeOfPrevPart != nil {
			lastNodeOfPrevPart.next = prev
		} else {
			head = prev
		}

		lastNodeOfSubList.next = current
		if current == nil {
			break
		}
		prev = lastNodeOfSubList
	}
	return head

}
func printMyNode2(head *MyNode2) {
	current := head
	for current != nil {
		fmt.Print(current.value)
		fmt.Print(" ")
		current = current.next
	}
	fmt.Println()
}

func main() {
	head := &MyNode2{value: 1}
	head.next = &MyNode2{value: 2}
	head.next.next = &MyNode2{value: 3}
	head.next.next.next = &MyNode2{value: 4}
	head.next.next.next.next = &MyNode2{value: 5}
	head.next.next.next.next.next = &MyNode2{value: 6}
	head.next.next.next.next.next.next = &MyNode2{value: 7}
	head.next.next.next.next.next.next.next = &MyNode2{value: 8}
	head = reverseEveryKElements(head, 3)
	printMyNode2(head)

}
