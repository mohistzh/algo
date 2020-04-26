package main

import (
	"fmt"

	"github.com/cheekybits/genny/generic"
)

// Node doublely LinkedList Node
type Node struct {
	Key, Value generic.Type
	Next, Prev *Node
}

// LRUCache structure
type LRUCache struct {
	Capacity   int
	Map        map[generic.Type]*Node
	Head, Tail *Node
}

// Constructor init LRUCache
func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		Capacity: capacity,
		Map:      make(map[generic.Type]*Node, capacity),
		Head:     &Node{},
		Tail:     &Node{},
	}
	lruCache.Head.Next = lruCache.Tail
	lruCache.Tail.Prev = lruCache.Head
	return lruCache
}

// Get get
func (lruCache *LRUCache) Get(key generic.Type) generic.Type {
	node, ok := lruCache.Map[key]
	if !ok {
		return nil
	}
	lruCache.removeNode(node)
	lruCache.setHeadNode(node)
	return node
}
func (lruCache *LRUCache) removeNode(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}
func (lruCache *LRUCache) setHeadNode(node *Node) {
	lruCache.Head.Prev = node
	node.Next = lruCache.Head.Next
	lruCache.Head.Next = node
	node.Prev = lruCache.Head
}

// Put add new data
func (lruCache *LRUCache) Put(key generic.Type, value generic.Type) {
	node, ok := lruCache.Map[key]
	if ok {
		lruCache.removeNode(node)
	} else {
		if len(lruCache.Map) == lruCache.Capacity {
			delete(lruCache.Map, lruCache.Tail.Prev.Key)
			lruCache.removeNode(lruCache.Tail.Prev)
		}
		node = &Node{Key: key, Value: value}
		lruCache.Map[node.Key] = node
	}
	node.Value = value
	lruCache.setHeadNode(node)
}

// Print print node list
func (lruCache *LRUCache) Print() {
	head := lruCache.Head
	for head != nil && head.Next != nil {
		fmt.Print("key:", head.Key, " value: ", head.Value)
		head = head.Next
	}
}

func main() {
	lru := Constructor(5)
	lru.Put("hello", "world")
	lru.Put(2, 3)
	lru.Put(2.2, 3.3)
	lru.Put(4, "test")
	lru.Print()
	// fmt.Println(lru.Get("Get"))
	// fmt.Println(lru.Get("hello"))
	// fmt.Println(lru.Get("Get"))
}
