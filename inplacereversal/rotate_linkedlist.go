package main

import "fmt"

// MyNode4 singly linked node
type MyNode4 struct {
	value int
	next  *MyNode4
}

/*
Given the head of a Singly LinkedList and a number ‘k’, rotate the LinkedList to the right by ‘k’ nodes.

the solution is th define the roration is to take the sub-list of 'k' ending nodes of the linkedlist and connect them to the beginning
e.g.

1->2->3->4->5->6  k=3
the sub-list of 'k' ending nodes of the linkedlist is 4, 5, 6

We need to consider three more things:

1. Connect the last node of the LinkedList to the head, because the list will have a different tail after the rotation.
2. The new head of the LinkedList will be the node at the beginning of the sub-list. e.g. 4
3. The node right before the start of sub-list will be the new tail of the rotated LinkedList.

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
	// move to the sublist
	for i := 1; i < skipLength; i++ {
		lastNodeOfRotatedList = lastNodeOfRotatedList.next
	}
	// points to the sub-list of 'k' endoing nodes
	head = lastNodeOfRotatedList.next
	lastNodeOfRotatedList.next = nil // cut down the cyclic link
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
