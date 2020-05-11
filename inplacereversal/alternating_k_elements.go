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
	current := head
	var prev *MyNode3
	for {
		// we consider multip parts of the linked list, prev and current nodes used to link adjancy sub-list
		lastNodeOfPrevSubList := prev
		lastNodeOfCurSubList := current
		count := 0
		for current != nil && count < k {
			next := current.next
			current.next = prev
			prev = current
			current = next
			count++
		}
		if lastNodeOfPrevSubList != nil {
			lastNodeOfPrevSubList.next = prev
		} else {
			head = prev
		}
		lastNodeOfCurSubList.next = current
		count = 0
		// skipping k elements
		for current != nil && count < k {
			prev = current
			current = current.next
			count++
		}
		if current == nil {
			break
		}

	}
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
	head.next.next.next.next.next.next.next.next = &MyNode3{value: 9}
	head.next.next.next.next.next.next.next.next.next = &MyNode3{value: 10}
	head = reverseAlternatingKElements(head, 3)
	printMyNode3(head)

}
