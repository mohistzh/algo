package linkedlist

import (
	"fmt"
	"sync"

	"github.com/cheekybits/genny/generic"
)

// Item type of the linked list
type Item generic.Type

// Node a single node that composes the list
type Node struct {
	content Item
	next    *Node
}

// SinglyLinkedList the linked list of Items
type SinglyLinkedList struct {
	head *Node
	size int
	lock sync.RWMutex
}

// Append adds an Item to the end of the linked list
func (list *SinglyLinkedList) Append(item Item) {
	list.lock.Lock()
	node := Node{item, nil}
	if list.head == nil {
		list.head = &node
	} else { // append item to the end
		last := list.head
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = &node
	}
	list.size++
	list.lock.Unlock()
}

// Insert adds an Item at position index
func (list *SinglyLinkedList) Insert(index int, item Item) error {
	list.lock.Lock()
	defer list.lock.Unlock()
	if index < 0 || index > list.size {
		return fmt.Errorf("Index out fo bounds")
	}
	newNode := Node{item, nil}
	if index == 0 { // append new node at the beginning
		newNode.next = list.head
		list.head = &newNode
		return nil
	}
	node := list.head
	pos := 0
	for pos < index-2 {
		pos++
		node = node.next
	}
	newNode.next = node.next
	node.next = &newNode
	list.size++
	return nil
}

//RemoveAt removes a node at position
func (list *SinglyLinkedList) RemoveAt(index int) (*Item, error) {
	list.lock.Lock()
	defer list.lock.Lock()
	if index < 0 || index > list.size {
		return nil, fmt.Errorf("Index out of bounds")
	}
	node := list.head
	pos := 0
	for pos < index-1 { // to find node which need to be deleted
		pos++
		node = node.next
	}
	removedNode := node.next
	node.next = removedNode.next
	list.size--
	return &removedNode.content, nil
}

// IndexOf return the position of the Item
func (list *SinglyLinkedList) IndexOf(item Item) int {
	list.lock.RLock()
	defer list.lock.RUnlock()
	node := list.head
	pos := 0
	for {
		if node.content == item {
			return pos
		}
		if node.next == nil {
			return -1
		}
		node = node.next
		pos++
	}
}

// IsEmpty return true if the list is empty
func (list *SinglyLinkedList) IsEmpty() bool {
	list.lock.RLock()
	defer list.lock.RUnlock()
	if list.head == nil {
		return true
	}
	return false
}

// Size return the linked list size
func (list *SinglyLinkedList) Size() int {
	list.lock.RLock()
	defer list.lock.RUnlock()
	size := 1
	last := list.head
	for {
		if last == nil || last.next == nil {
			break
		}
		last = last.next
		size++
	}
	return size
}

// String print linked list
func (list *SinglyLinkedList) String() {
	list.lock.Lock()
	defer list.lock.Unlock()
	node := list.head
	pos := 0
	for {
		if node == nil {
			break
		}
		pos++
		fmt.Print(node.content)
		fmt.Print(" ")
		node = node.next
	}
	fmt.Println()
}

// Head return the head of the linked list
func (list *SinglyLinkedList) Head() *Node {
	list.lock.RLock()
	defer list.lock.RUnlock()
	return list.head
}
