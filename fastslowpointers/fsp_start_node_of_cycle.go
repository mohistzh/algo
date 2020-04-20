package main

import "fmt"

type SinglyNode struct {
	value int
	next  *SinglyNode
}

/*
Given the head of a Singly LinkedList that contains a cycle, write a function to find the starting node of the cycle.
*/
func startOfLinkedListCycle(head *SinglyNode) *SinglyNode {
	slow, fast := head, head.next
	lengthOfCycle := 0
	for fast != nil && fast.next != nil {
		if slow.value == fast.value {
			lengthOfCycle = findCycleLength(slow)
			break
		}
		slow = slow.next
		fast = fast.next.next
	}
	return findCycleStartNode(head, lengthOfCycle)
}

func findCycleStartNode(head *SinglyNode, k int) *SinglyNode {
	pointer1, pointer2 := head, head
	for k > 0 {
		k--
		pointer2 = pointer2.next // move pointer2 ahead of pointer1 'k' nodes
	}
	for pointer1 != pointer2 {
		pointer1 = pointer1.next
		pointer2 = pointer2.next
	}
	return pointer1
}

func findCycleLength(head *SinglyNode) int {
	current := head
	length := 0
	for {
		current = current.next
		length++
		if current == head {
			return length
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
	head.next.next.next.next.next.next = head.next.next
	fmt.Println("LinkedList start node of cycle:", startOfLinkedListCycle(&head).value) // 3
	head.next.next.next.next.next.next = head.next.next.next
	fmt.Println("LinkedList start node of cycle:", startOfLinkedListCycle(&head).value) // 4
	head.next.next.next.next.next.next = &head
	fmt.Println("LinkedList start node of cycle:", startOfLinkedListCycle(&head).value) // 1
}
