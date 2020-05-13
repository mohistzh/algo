package main

import (
	"fmt"
)

// MyTreeNode5 tree node
type MyTreeNode5 struct {
	value       int
	left, right *MyTreeNode5
}

/*
Find the minimum depth of a binary tree. The minimum depth is the number of nodes along the shortest path from the root node to the nearest leaf node.

*/
func findMinimumDepth(root *MyTreeNode5) int {
	minimum := 0
	var queue []*MyTreeNode5
	queue = append(queue, root)
	for len(queue) > 0 { // tree level
		levelSize := len(queue)
		minimum++
		for i := 0; i < levelSize; i++ { // node size
			node := queue[0]
			queue = queue[1:]
			if node.left == nil && node.right == nil {
				return minimum
			}
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
	}
	return minimum
}

func main() {
	root := &MyTreeNode5{value: 12}
	root.left = &MyTreeNode5{value: 7}
	root.right = &MyTreeNode5{value: 1}
	root.left.left = &MyTreeNode5{value: 9}
	root.right.left = &MyTreeNode5{value: 10}
	root.right.right = &MyTreeNode5{value: 5}
	root.right.left.left = &MyTreeNode5{value: 11}
	fmt.Println(findMinimumDepth(root))
}
