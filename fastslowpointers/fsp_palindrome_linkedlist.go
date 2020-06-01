package main

import (
	"fmt"
)

// SinglyNode basic node structure
type SinglyNode struct {
	value int
	next  *SinglyNode
}

/*
Given the head of a Singly LinkedList, write a method to check if the LinkedList is a palindrome or not.

Your algorithm should use constant space and the input LinkedList should be in the original form once the algorithm is finished.
The algorithm should have O(N)O(N) time complexity where ‘N’ is the number of nodes in the LinkedList.
*/
func findPalindrome(head *SinglyNode) bool {
	slow, fast, curHead := head, head, head
	for fast != nil && fast.next != nil { // use slow and fast pointers to reach the middle of the linkedlist
		slow = slow.next
		fast = fast.next.next
	}
	headSecondHalf := reverse(slow) // reverse the second half, comparing with the first half
	copyOfHeadSecondHalf := headSecondHalf
	for curHead != nil && headSecondHalf != nil {
		if curHead.value != headSecondHalf.value {
			break
		}
		curHead = curHead.next
		headSecondHalf = headSecondHalf.next
	}
	reverse(copyOfHeadSecondHalf)
	if headSecondHalf == nil {
		return true
	}
	return false
}

func reverse(head *SinglyNode) *SinglyNode {
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
	head := SinglyNode{value: 2}
	head.next = &SinglyNode{value: 4}
	head.next.next = &SinglyNode{value: 6}
	head.next.next.next = &SinglyNode{value: 4}
	head.next.next.next.next = &SinglyNode{value: 2}
	fmt.Println(findPalindrome(&head))
	head.next.next.next.next.next = &SinglyNode{value: 2}
	fmt.Println(findPalindrome(&head))
	head = SinglyNode{value: 0}
	head.next = &SinglyNode{value: 0}
	fmt.Println(findPalindrome(&head))
}
