package main

import (
	"fmt"
)

// MyTreeNode6 tree node
type MyTreeNode6 struct {
	value       int
	left, right *MyTreeNode6
}

/*
Given a binary tree, find its maximum depth (or height).
*/
func findMaximumDepth(root *MyTreeNode6) int {
	maximum := 0
	var queue []*MyTreeNode6
	queue = append(queue, root)
	for len(queue) > 0 {
		maximum++
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
	}
	return maximum

}

func main() {
	root := &MyTreeNode6{value: 12}
	root.left = &MyTreeNode6{value: 7}
	root.right = &MyTreeNode6{value: 1}
	root.left.left = &MyTreeNode6{value: 9}
	root.right.left = &MyTreeNode6{value: 10}
	root.right.right = &MyTreeNode6{value: 5}
	root.right.left.left = &MyTreeNode6{value: 11}

	fmt.Println(findMaximumDepth(root))

}
