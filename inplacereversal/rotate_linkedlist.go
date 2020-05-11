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
	head.next.next.next.next.next = &MyNode4{value: 6}
	head = rotateLinkedList(head, 3)
	printMyNode4(head)
}
