package main

import (
	"fmt"
)

type SinglyNode struct {
	value int
	next  *SinglyNode
}

/*
Given the head of a Singly LinkedList, write a method to modify the LinkedList such that the nodes from the second half of the LinkedList
are inserted alternately to the nodes from the first half in reverse order.
So if the LinkedList has nodes 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> null, your method should return 1 -> 6 -> 2 -> 5 -> 3 -> 4 -> null.
e.g.
Input: 2 -> 4 -> 6 -> 8 -> 10 -> 12 -> null
Output: 2 -> 12 -> 4 -> 10 -> 6 -> 8 -> null
*/
func rearrangeLinkedList(head *SinglyNode) {
	slow, fast := head, head
	for fast != nil && fast.next != nil { // found the middle
		slow = slow.next
		fast = fast.next.next
	}
	secondHalfHead := reverseSecondHalf(slow) // reverse the second half of the linked list
	firstHalfHead := head

	for firstHalfHead != nil && secondHalfHead != nil {
		tmp := firstHalfHead.next
		firstHalfHead.next = secondHalfHead
		firstHalfHead = tmp

		tmp = secondHalfHead.next
		secondHalfHead.next = firstHalfHead
		secondHalfHead = tmp
	}
	if firstHalfHead != nil {
		firstHalfHead.next = nil
	}
	temp := head
	for temp != nil {
		fmt.Print(temp.value, " -> ")
		temp = temp.next
	}
}

func reverseSecondHalf(head *SinglyNode) *SinglyNode {
	var prev *SinglyNode
	for head != nil {
		next := head.next
		head.next = prev
		prev = head
		head = next
	}
	return prev
}

func main() {
	head := &SinglyNode{value: 2}
	head.next = &SinglyNode{value: 4}
	head.next.next = &SinglyNode{value: 6}
	head.next.next.next = &SinglyNode{value: 8}
	head.next.next.next.next = &SinglyNode{value: 10}
	head.next.next.next.next.next = &SinglyNode{value: 12}
	rearrangeLinkedList(head)
}
