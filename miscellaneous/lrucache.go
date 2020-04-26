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
		Store:    make(map[generic.Type]*Node, capacity),
		Head:     &Node{},
		Tail:     &Node{},
	}
	// default states
	lruCache.Head.Next = lruCache.Tail
	lruCache.Tail.Prev = lruCache.Head

	return lruCache
}

// Get get
func (lruCache *LRUCache) Get(key generic.Type) generic.Type {
	node, ok := lruCache.Store[key]
	if !ok {
		return nil
	}
	//refresh node position
	lruCache.removeNode(node)
	lruCache.setHead(node)

	return node
}

// removeNode remove the given node
func (lruCache *LRUCache) removeNode(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

// setHead reset head node
func (lruCache *LRUCache) setHead(node *Node) {
	lruCache.Head.Prev = node
	node.Next = lruCache.Head.Next
	lruCache.Head.Next = node
	node.Prev = lruCache.Head
}

// ensureCapacity calling it when cache data is going to exceeded the lurcache capacity
func (lruCache *LRUCache) ensureCapacity() {
	if len(lruCache.Store) == lruCache.Capacity {
		delete(lruCache.Store, lruCache.Tail.Prev.Key) // Delete the last item of the Store
		newLast := lruCache.Tail.Prev.Prev
		lruCache.removeNode(lruCache.Tail.Prev)
		lruCache.Tail.Prev = newLast
		newLast.Next = lruCache.Tail
	}
}

// Put add new data
func (lruCache *LRUCache) Put(key generic.Type, value generic.Type) {
	node, ok := lruCache.Store[key]
	if ok {
		lruCache.removeNode(node)
		node.Value = value
	} else {
		lruCache.ensureCapacity()
		node = &Node{Key: key, Value: value}
		lruCache.Store[key] = node
	}
	lruCache.setHead(node)

}

// Print print node list
func (lruCache *LRUCache) Print() {
	head := lruCache.Head.Next
	for head != nil && head.Next != nil {
		fmt.Print("key:", head.Key, " value:", head.Value)
		fmt.Println()
		head = head.Next
	}
	fmt.Println("---------------")
}

func main() {
	lru := Constructor(5)
	lru.Put("hello", "world")
	lru.Put(2, 3)
	lru.Put(2.2, 3.3)
	lru.Put(4, "test")
	lru.Put(9, "888")
	lru.Print()
	fmt.Println(lru.Get("Get"))
	fmt.Println(lru.Get("hello"))
	fmt.Println(lru.Get(4))
	lru.Put(8, 88)
}
