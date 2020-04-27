package main

import (
	"fmt"

	"github.com/cheekybits/genny/generic"
)

/*
Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.

get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity,
it should invalidate the least recently used item before inserting a new item.

The cache is initialized with a positive capacity.
*/

// Node doublely linked list node
type Node struct {
	Key, Value generic.Type
	Prev, Next *Node
}

// LRUCache the structure consists of LinkedList and Map
type LRUCache struct {
	Capacity   int
	Store      map[generic.Type]*Node // key links to Node (k & v)
	Head, Tail *Node                  // head and tail are references for links into Store nodes, they don't store actual data
}

// Constructor init LRUCache
func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		Capacity: capacity,
		Store:    make(map[generic.Type]*Node),
		Head:     &Node{},
		Tail:     &Node{},
	}
	lruCache.Head.Next = lruCache.Tail
	lruCache.Tail.Prev = lruCache.Head
	lruCache.Head.Prev = nil
	lruCache.Tail.Next = nil
	return lruCache
}

// Get get
func (lruCache *LRUCache) Get(key generic.Type) generic.Type {
	node, ok := lruCache.Store[key]
	if !ok {
		return nil
	}
	//refresh node position
	ans := node.Value
	lruCache.removeNode(node)
	lruCache.setHead(node)

	return ans
}

// Put add new data
func (lruCache *LRUCache) Put(key generic.Type, value generic.Type) {
	node, ok := lruCache.Store[key]
	if ok {
		node.Value = value
		lruCache.removeNode(node)
		lruCache.setHead(node)
	} else {
		node = &Node{Key: key, Value: value}
		lruCache.Store[key] = node
		if len(lruCache.Store) < lruCache.Capacity {
			lruCache.setHead(node)
		} else {
			delete(lruCache.Store, lruCache.Tail.Prev.Key)
			lruCache.removeNode(lruCache.Tail.Prev)
			lruCache.setHead(node)
		}

	}
}

// removeNode remove the given node
func (lruCache *LRUCache) removeNode(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev

}

// setHead reset head node
func (lruCache *LRUCache) setHead(node *Node) {
	if lruCache.Head == nil {
		lruCache.Tail = node
	} else {
		lruCache.Head.Prev = node
		node.Next = lruCache.Head
		node.Prev = nil
	}
	lruCache.Head = node

}

// Print print node list
func (lruCache *LRUCache) Print() {
	fmt.Println("---------------")
	fmt.Println("Store size:", len(lruCache.Store), " Store capacity:", lruCache.Capacity)
	fmt.Println(lruCache.Store)
	head := lruCache.Head
	for head != nil && head.Key != nil {
		fmt.Print(head.Key, " -> ", head.Value)
		fmt.Println()
		head = head.Next
	}
	fmt.Println("---------------")
}

func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1), " ")
	lru.Put(3, 3)
	fmt.Println(lru.Get(2), " ")
	lru.Put(4, 4)
	fmt.Println(lru.Get(1), " ")
	fmt.Println(lru.Get(3), " ")

	fmt.Println(lru.Get(4), " ")
	lru.Print()
}
