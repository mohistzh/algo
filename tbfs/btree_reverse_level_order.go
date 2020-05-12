package main

import (
	"fmt"
)

// MyTreeNode2 tree node
type MyTreeNode2 struct {
	value       int
	left, right *MyTreeNode2
}

/*
Given a binary tree, populate an array to represent its level-by-level traversal in reverse order,
i.e., the lowest level comes first. You should populate the values of all nodes in each level from left to right in separate sub-arrays.
*/
func reverseLevelOrderTraversal(root *MyTreeNode2) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	var queue []*MyTreeNode2
	queue = append(queue, root)
	for len(queue) > 0 {
		levelLength := len(queue)
		var levelNodes []int
		for i := 0; i < levelLength; i++ {
			node := queue[0]
			queue = queue[1:] // dequeu
			levelNodes = append(levelNodes, node.value)
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
		result = append([][]int{levelNodes}, result...) // prepend level nodes to the beginning.
	}
	return result

}

func main() {
	root := &MyTreeNode2{value: 1}
	root.left = &MyTreeNode2{value: 2}
	root.right = &MyTreeNode2{value: 3}
	root.left.left = &MyTreeNode2{value: 4}
	root.left.right = &MyTreeNode2{value: 5}
	root.right.left = &MyTreeNode2{value: 6}
	root.right.right = &MyTreeNode2{value: 7}
	fmt.Println(reverseLevelOrderTraversal(root))
}
