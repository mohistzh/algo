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
Given a binary tree, populate an array to represent its zigzag level order traversal. You should populate the values of all nodes of the first level from left to right,
then right to left for the next level and keep alternating in the same manner for the following levels.
*/
func zigzagTraversal(root *MyTreeNode3) [][]int {
	var result [][]int
	var queue []*MyTreeNode3 // to save downside level nodes
	zigzag := false          // root node state
	queue = append(queue, root)
	for len(queue) > 0 {
		levelLength := len(queue)
		var currentLevelQueue []int // to save current level node value with zigzag state order
		for i := 0; i < levelLength; i++ {
			currentNode := queue[0] // pop up
			queue = queue[1:]       // deque
			if zigzag {             // left -> right
				currentLevelQueue = append([]int{currentNode.value}, currentLevelQueue...)
			} else {
				currentLevelQueue = append(currentLevelQueue, currentNode.value)
			}

			if currentNode.left != nil {
				queue = append(queue, currentNode.left)
			}
			if currentNode.right != nil {
				queue = append(queue, currentNode.right)
			}
		}
		result = append(result, currentLevelQueue)
		zigzag = !zigzag // reverse the traversal direction

	}
	return result
}

func main() {
	root := &MyTreeNode3{value: 12}
	root.left = &MyTreeNode3{value: 7}
	root.right = &MyTreeNode3{value: 1}
	root.left.left = &MyTreeNode3{value: 9}
	root.right.left = &MyTreeNode3{value: 10}
	root.right.right = &MyTreeNode3{value: 5}
	root.right.left.left = &MyTreeNode3{value: 20}
	root.right.right.right = &MyTreeNode3{value: 17}

	fmt.Println(zigzagTraversal(root))
}
