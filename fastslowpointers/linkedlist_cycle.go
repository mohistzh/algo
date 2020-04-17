package main

import (
	"github.com/mohistzh/algo/linkedlist"
)

func hasCycle(head *linkedlist.Node) bool {
	return false
}

func main() {
	var singlyLinkedList linkedlist.SinglyLinkedList
	singlyLinkedList.Append(1)
	singlyLinkedList.Append(2)
	singlyLinkedList.Append(3)
	singlyLinkedList.Append(4)
	singlyLinkedList.Append(5)
	singlyLinkedList.Append(6)
	singlyLinkedList.String()
	singlyLinkedList.Append(4)
	singlyLinkedList.String()

}
