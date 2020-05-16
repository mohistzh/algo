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
Given a binary tree and a number ‘S’, find if the tree has a path from root-to-leaf such that the sum of all the node values of that path equals ‘S’.

*/
func findBinaryTreePathWithSum(root *MyTreeNode1, sum int) bool {
	if root == nil {
		return false
	}
	if root.value == sum && root.left == nil && root.right == nil {
		return true
	}
	return findBinaryTreePathWithSum(root.left, sum-root.value) || findBinaryTreePathWithSum(root.right, sum-root.value)
}

// traversal solution with stack TODO
func findBinaryTreePathWithSum2(root *MyTreeNode1, sum int) bool {
	if root == nil {
		return false
	}
	var stack []*MyTreeNode1
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Print(node.value, " ")
		if node.right != nil {
			stack = append(stack, node.right)

		}
		if node.left != nil {
			stack = append(stack, node.left)
		}
	}
	return false
}

func main() {
	root := &MyTreeNode1{value: 12}
	root.left = &MyTreeNode1{value: 7}
	root.right = &MyTreeNode1{value: 1}
	root.left.left = &MyTreeNode1{value: 9}
	root.right.left = &MyTreeNode1{value: 10}
	root.right.right = &MyTreeNode1{value: 5}

	fmt.Println(findBinaryTreePathWithSum(root, 23))
	fmt.Println(findBinaryTreePathWithSum(root, 16))

	fmt.Println(findBinaryTreePathWithSum2(root, 23))
}
