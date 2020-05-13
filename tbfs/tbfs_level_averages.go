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
Given a binary tree, populate an array to represent the averages of all of its levels.
*/
func levelAverages(root *MyTreeNode4) []int {
	return nil
}

func main() {
	root := &MyTreeNode4{value: 1}
	root.left = &MyTreeNode4{value: 2}
	root.right = &MyTreeNode4{value: 3}
	root.left.left = &MyTreeNode4{value: 4}
	root.left.right = &MyTreeNode4{value: 5}
	root.right.left = &MyTreeNode4{value: 6}
	root.right.right = &MyTreeNode4{value: 7}

	fmt.Println(levelAverages(root))
}
