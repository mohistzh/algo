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
	if root == nil {
		return result
	}
	// use a queue to store nodes of each level
	var queue []*MyTreeNode1
	queue = append(queue, root)
	for len(queue) > 0 {
		levelSize := len(queue)
		var currentLevel []int // use to store value of nodes
		for i := 0; i < levelSize; i++ {
			currentNode := queue[0]
			queue = queue[1:] // dequeue top element
			currentLevel = append(currentLevel, currentNode.value)
			if currentNode.left != nil {
				queue = append(queue, currentNode.left) // fill in left node
			}
			if currentNode.right != nil {
				queue = append(queue, currentNode.right) // fill in right node
			}
		}
		result = append(result, currentLevel) // store level
	}
	return result
}

/*
	12
	7 1
	9 10 5
*/
func main() {
	root := &MyTreeNode1{value: 12}
	root.left = &MyTreeNode1{value: 7}
	root.right = &MyTreeNode1{value: 1}
	root.left.left = &MyTreeNode1{value: 9}
	root.right.left = &MyTreeNode1{value: 10}
	root.right.right = &MyTreeNode1{value: 5}
	fmt.Println("Level order traversal:", traverseMyTreeNode1(root))

}
