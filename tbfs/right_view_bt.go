package main

import (
	"fmt"
)

// MyTreeNode10 tree node
type MyTreeNode10 struct {
	value       int
	left, right *MyTreeNode10
}

/*
Given a binary tree, return an array containing nodes in its right view.
The right view of a binary tree is the set of nodes visible when the tree is seen from the right side.
*/
func rightViewBinaryTree(root *MyTreeNode10) []int {
	var result []int

	return result
}

func main() {
	root := &MyTreeNode10{value: 12}
	root.left = &MyTreeNode10{value: 7}
	root.right = &MyTreeNode10{value: 1}
	root.left.left = &MyTreeNode10{value: 9}
	root.right.left = &MyTreeNode10{value: 10}
	root.right.right = &MyTreeNode10{value: 5}
	root.left.left = &MyTreeNode10{value: 3}
	fmt.Println(rightViewBinaryTree(root))
}
