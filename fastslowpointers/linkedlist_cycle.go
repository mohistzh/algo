package main

import (
	"fmt"
)

// SinglyNode basic node structure
type SinglyNode struct {
	value int
	next  *SinglyNode
}

// hasCycle super easy question, fast and slow pointers pattern
func hasCycle(head *SinglyNode) bool {
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
	head := SinglyNode{value: 1}
	head.next = &SinglyNode{value: 2}
	head.next.next = &SinglyNode{value: 3}
	head.next.next.next = &SinglyNode{value: 4}
	head.next.next.next.next = &SinglyNode{value: 5}
	head.next.next.next.next.next = &SinglyNode{value: 6}
	fmt.Println("LinkedList has cycle:", hasCycle(&head))
	head.next.next.next.next.next.next = head.next.next
	fmt.Println("LinkedList has cycle:", hasCycle(&head))
	head.next.next.next.next.next.next = head.next.next.next
	fmt.Println("LinkedList has cycle:", hasCycle(&head))
}
