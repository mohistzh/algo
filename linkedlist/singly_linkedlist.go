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
