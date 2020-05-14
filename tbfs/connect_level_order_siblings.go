package main

import (
	"fmt"
)

// MyTreeNode8 tree node
type MyTreeNode8 struct {
	value             int
	left, right, next *MyTreeNode8
}

/*
Given a binary tree, connect each node with its level order successor. The last node of each level should point to a null node.
*/
func connectLevelOrderSiblings(root *MyTreeNode8) *MyTreeNode8 {
	var queue []*MyTreeNode8
	queue = append(queue, root)

	for len(queue) > 0 {
		var prev *MyTreeNode8
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if prev != nil {
				prev.next = node
			}
			prev = node

			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}

	}
	return root
}

func printMyTreeNode8(root *MyTreeNode8) {
	nextLevelRoot := root
	for nextLevelRoot != nil {
		current := nextLevelRoot
		nextLevelRoot = nil
		for current != nil {
			fmt.Print(current.value, " ")
			if nextLevelRoot == nil {
				if current.left != nil {
					nextLevelRoot = current.left
				} else if current.right != nil {
					nextLevelRoot = current.right
				}
			}
			current = current.next
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
	root = connectLevelOrderSiblings(root)
	printMyTreeNode8(root)
}
