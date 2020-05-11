package main

import "fmt"

// MyNode4 singly linked node
type MyNode4 struct {
	value int
	next  *MyNode4
}

/*
Given the head of a Singly LinkedList and a number ‘k’, rotate the LinkedList to the right by ‘k’ nodes.
*/
func rotateLinkedList(head *MyNode4, k int) *MyNode4 {
	lastNode := head
	listLength := 1
	// count length of the linked list and get the last node
	for lastNode.next != nil {
		lastNode = lastNode.next
		listLength++
	}
	lastNode.next = head         // make a circular linked list
	k %= listLength              // in case the k is great than the list size, we don't need to rotate all length
	skipLength := listLength - k //
	lastNodeOfRotatedList := head
	for i := 1; i < skipLength; i++ {
		lastNodeOfRotatedList = lastNodeOfRotatedList.next
	}
	// points to the sub-list of 'k' endoing nodes
	head = lastNodeOfRotatedList.next
	lastNodeOfRotatedList.next = nil
	return head

}
func printMyNode4(head *MyNode4) {
	current := head
	for current != nil {
		fmt.Print(current.value)
		fmt.Print(" ")
		current = current.next
	}
	fmt.Println()
}

func main() {
	head := &MyNode4{value: 1}
	head.next = &MyNode4{value: 2}
	head.next.next = &MyNode4{value: 3}
	head.next.next.next = &MyNode4{value: 4}
	head.next.next.next.next = &MyNode4{value: 5}
	head = rotateLinkedList(head, 8)
	printMyNode4(head)
}
