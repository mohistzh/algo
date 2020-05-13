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
	var result []int
	var queue []*MyTreeNode4
	queue = append(queue, root)
	for len(queue) > 0 {
		var levelNodeQueue []int
		levelLength := len(queue)
		for i := 0; i < levelLength; i++ {
			node := queue[0]
			queue = queue[1:]
			levelNodeQueue = append(levelNodeQueue, node.value)
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
		length := len(levelNodeQueue)
		sum := 0
		for i := 0; i < length; i++ {
			sum += levelNodeQueue[i]
		}
		result = append(result, sum/length)
	}

	return result
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
