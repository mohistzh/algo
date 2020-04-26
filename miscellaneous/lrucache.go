package main

import (
	"fmt"

	"github.com/cheekybits/genny/generic"
)

// Node doublely LinkedList Node
type Node struct {
	Key, Value generic.Type
	next, prev *Node
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
	lruCache.Head.next = lruCache.Tail
	lruCache.Tail.prev = lruCache.Head
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
	node.prev.next = node.next
	node.next.prev = node.prev
}
func (lruCache *LRUCache) setHeadNode(node *Node) {
	lruCache.Head.prev = node
	node.next = lruCache.Head.next
	lruCache.Head.next = node
	node.prev = lruCache.Head
}

// Put add new data
func (lruCache *LRUCache) Put(key generic.Type, value generic.Type) {
	node, ok := lruCache.Map[key]
	if ok {
		lruCache.removeNode(node)
	} else {
		if len(lruCache.Map) == lruCache.Capacity {
			delete(lruCache.Map, lruCache.Tail.prev.Key)
			lruCache.removeNode(lruCache.Tail.prev)
		}
		node = &Node{Key: key, Value: value}
		lruCache.Map[node.Key] = node
	}
	node.Value = value
	lruCache.setHeadNode(node)
}
func (lruCache *LRUCache) Print() {
	head := lruCache.Head
	for head != nil && head.next != nil {
		fmt.Print("key:", head.Key, " value: ", head.Value)
		head = head.next
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
