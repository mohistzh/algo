package main

import (
	"fmt"
)

// MyTreeNode3 tree node
type MyTreeNode3 struct {
	value       int
	left, right *MyTreeNode3
}

/*
Given a binary tree where each node can only have a digit (0-9) value,
each root-to-leaf path will represent a number.
Find the total sum of all the numbers represented by all paths.
*/
func sumOfPathNumbers(root *MyTreeNode3) int {
	return -1
}

func main() {
	root := &MyTreeNode3{value: 1}
	root.left = &MyTreeNode3{value: 7}
	root.right = &MyTreeNode3{value: 9}
	root.right.left = &MyTreeNode3{value: 2}
	root.right.right = &MyTreeNode3{value: 9}
	fmt.Println(sumOfPathNumbers(root))
}
