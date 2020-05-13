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
	minimum := -1

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
