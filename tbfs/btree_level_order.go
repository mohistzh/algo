package main

import (
	"fmt"
)

// MyTreeNode1 tree node
type MyTreeNode1 struct {
	value       int
	left, right *MyTreeNode1
}

/*
Given a binary tree, populate an array to represent its level-by-level traversal.
You should populate the values of all nodes of each level from left to right in separate sub-arrays.
*/
func traverseMyTreeNode1(root *MyTreeNode1) [][]int {
	var result [][]int
	return result
}

func main() {
	root := &MyTreeNode1{value: 12}
	root.left = &MyTreeNode1{value: 7}
	root.right = &MyTreeNode1{value: 1}
	root.left.left = &MyTreeNode1{value: 9}
	root.right.left = &MyTreeNode1{value: 10}
	root.right.right = &MyTreeNode1{value: 5}
	fmt.Println("Level order traversal:", traverseMyTreeNode1(root))

}
