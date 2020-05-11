package main

import "fmt"

// MyNode3 singly linked node
type MyNode3 struct {
	value int
	next  *MyNode3
}

/*
Given the head of a LinkedList and a number ‘k’, reverse every alternating ‘k’ sized sub-list starting from the head.

If, in the end, you are left with a sub-list with less than ‘k’ elements, reverse it too.
*/
func reverseAlternatingKElements(head *MyNode3, k int) *MyNode3 {

	return head
}

func printMyNode3(head *MyNode3) {
	current := head
	for current != nil {
		fmt.Print(current.value)
		fmt.Print(" ")
		current = current.next
	}
	fmt.Println()
}

func main() {
	head := &MyNode3{value: 1}
	head.next = &MyNode3{value: 2}
	head.next.next = &MyNode3{value: 3}
	head.next.next.next = &MyNode3{value: 4}
	head.next.next.next.next = &MyNode3{value: 5}
	head.next.next.next.next.next = &MyNode3{value: 6}
	head.next.next.next.next.next.next = &MyNode3{value: 7}
	head.next.next.next.next.next.next.next = &MyNode3{value: 8}
	head = reverseAlternatingKElements(head, 2)
	printMyNode3(head)

}
