package main

import (
	"fmt"
	"math"
)

// MyTreeNode6 tree node
type MyTreeNode6 struct {
	value       int
	left, right *MyTreeNode6
}

/*
Given a binary tree, find the length of its diameter. The diameter of a tree is the number of nodes on the longest path between any two leaf nodes.
The diameter of a tree may or may not pass through the root.

Note: You can always assume that there are at least two leaf nodes in the given tree.
*/
var treeDiameter = 0

func findTreeDiameter(root *MyTreeNode6) int {
	calculateHeight(root)
	return treeDiameter
}
func calculateHeight(currentNode *MyTreeNode6) int {
	if currentNode == nil {
		return 0
	}
	leftHeight := calculateHeight(currentNode.left)
	rightHeight := calculateHeight(currentNode.right)
	diameter := leftHeight + rightHeight + 1
	treeDiameter = int(math.Max(float64(diameter), float64(treeDiameter)))
	return int(math.Max(float64(leftHeight), float64(rightHeight))) + 1
}

func main() {
	root := &MyTreeNode6{value: 1}
	root.left = &MyTreeNode6{value: 2}
	root.right = &MyTreeNode6{value: 3}
	root.left.left = &MyTreeNode6{value: 4}
	root.right.left = &MyTreeNode6{value: 5}
	root.right.right = &MyTreeNode6{value: 6}
	fmt.Println("Tree diameter:", findTreeDiameter(root))
	root.left.left = nil
	root.right.left.left = &MyTreeNode6{value: 7}
	root.right.left.right = &MyTreeNode6{value: 8}
	root.right.right = &MyTreeNode6{value: 9}
	root.right.left.right.right = &MyTreeNode6{value: 10}
	root.right.right = &MyTreeNode6{value: 11}
	fmt.Println("Tree diameter:", findTreeDiameter(root))
}
