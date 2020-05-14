package main

import (
	"fmt"
)

// MyTreeNode8 tree node
type MyTreeNode8 struct {
	value       int
	left, right *MyTreeNode8
}

/*
Given a binary tree, connect each node with its level order successor. The last node of each level should point to a null node.
*/
func connectLevelOrderSiblings(root *MyTreeNode8) *MyTreeNode8 {
	printMyTreeNode8(root)
	return nil
}

func printMyTreeNode8(root *MyTreeNode8) {
	var queue []*MyTreeNode8
	queue = append(queue, root)
	for len(queue) > 0 {
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
			fmt.Print(node.value)
			fmt.Print(" ")
		}
		fmt.Println()

	}
}

func main() {
	root := &MyTreeNode8{value: 1}
	root.left = &MyTreeNode8{value: 2}
	root.right = &MyTreeNode8{value: 3}
	root.left.left = &MyTreeNode8{value: 4}
	root.left.right = &MyTreeNode8{value: 5}
	root.right.left = &MyTreeNode8{value: 6}
	root.right.right = &MyTreeNode8{value: 7}
	fmt.Println("Before")
	printMyTreeNode8(root)
	fmt.Println("After")
	connectLevelOrderSiblings(root)
}
