package main

import (
	"fmt"
)

// MyTreeNode4 tree node
type MyTreeNode4 struct {
	value       int
	left, right *MyTreeNode4
}

/*
Given a binary tree and a number sequence, find if the sequence is present as a root-to-leaf path in the given tree.
*/
func pathWithGivenSequence(root *MyTreeNode4, sequence []int) bool {
	return false
}

func main() {
	root := &MyTreeNode4{value: 1}
	root.left = &MyTreeNode4{value: 7}
	root.right = &MyTreeNode4{value: 9}
	root.right.left = &MyTreeNode4{value: 2}
	root.right.right = &MyTreeNode4{value: 9}

	fmt.Println(root, []int{1, 9, 9})
}
