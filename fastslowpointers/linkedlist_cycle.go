package main

import (
	"fmt"
)

//node linked list node
type node struct {
	value int
	next  *node
}

// hasCycle super easy question, fast and slow pointers pattern
func hasCycle(head *node) bool {
	slow := head
	fast := head.next
	for fast != nil && fast.next != nil {
		if slow.value == fast.value {
			return true
		}
		slow = slow.next
		fast = fast.next.next
	}
	return false
}

func main() {
	head := node{value: 1}
	head.next = &node{value: 2}
	head.next.next = &node{value: 3}
	head.next.next.next = &node{value: 4}
	head.next.next.next.next = &node{value: 5}
	head.next.next.next.next.next = &node{value: 6}
	fmt.Println("LinkedList has cycle:", hasCycle(&head))
	head.next.next.next.next.next.next = head.next.next
	fmt.Println("LinkedList has cycle:", hasCycle(&head))
	head.next.next.next.next.next.next = head.next.next.next
	fmt.Println("LinkedList has cycle:", hasCycle(&head))
}
