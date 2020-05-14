package main

import (
	"fmt"
)

// MyTreeNode7 tree node
type MyTreeNode7 struct {
	value       int
	left, right *MyTreeNode7
}

/*
Given a binary tree and a node, find the level order successor of the given node in the tree.
The level order successor is the node that appears right after the given node in the level order traversal.
*/
func findLevelOrderSuccessor(root *MyTreeNode7, target int) *MyTreeNode7 {
	var queue []*MyTreeNode7
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}
		if node.value == target {
			break
		}
	}
	return queue[0]
}

func main() {
	root := &MyTreeNode7{value: 1}
	root.left = &MyTreeNode7{value: 2}
	root.right = &MyTreeNode7{value: 3}
	root.left.left = &MyTreeNode7{value: 4}
	root.left.left = &MyTreeNode7{value: 5}
	fmt.Println(findLevelOrderSuccessor(root, root.right.value))
}
