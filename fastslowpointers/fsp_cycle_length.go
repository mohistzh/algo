package main

import (
	"fmt"
)

// SinglyNode basic node structure
type SinglyNode struct {
	value int
	next  *SinglyNode
}

// Given the head of a LinkedList with a cycle, find the length of the cycle.
func findLengthOfCycle(head *SinglyNode) int {
	slow, fast := head, head.next
	for fast != nil && fast.next != nil {
		if slow.value == fast.value {
			return findLength(slow)
		}
		slow = slow.next
		fast = fast.next.next
	}
	return -1
}

func findLength(slow *SinglyNode) int {
	current := slow
	length := 0
	for {
		current = current.next
		length++
		if current.value == slow.value {
			break
		}
	}
	return length
}

func main() {
	head := SinglyNode{value: 1}
	head.next = &SinglyNode{value: 2}
	head.next.next = &SinglyNode{value: 3}
	head.next.next.next = &SinglyNode{value: 4}
	head.next.next.next.next = &SinglyNode{value: 5}
	head.next.next.next.next.next = &SinglyNode{value: 6}
	fmt.Println("LinkedList cycle length:", findLengthOfCycle(&head))
	head.next.next.next.next.next.next = head.next.next
	fmt.Println("LinkedList cycle length:", findLengthOfCycle(&head))
	head.next.next.next.next.next.next = head.next.next.next
	fmt.Println("LinkedList cycle length:", findLengthOfCycle(&head))
}
