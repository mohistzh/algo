package main

import (
	"fmt"
)

// MyTreeNode3 tree node
type MyTreeNode3 struct {
	value       int
	left, right *MyTreeNode3
}

func zigzagTraversal(root *MyTreeNode3) [][]int {
	var result [][]int
	zigzag := false
	var queue []*MyTreeNode3
	queue = append(queue, root)

	for len(queue) > 0 {
		levelLength := len(queue)
		var currentLevel []int
		for i := 0; i < levelLength; i++ {
			node := queue[0]
			queue = queue[1:] // deque
			if zigzag {       // if zigzag is true, from right to left
				currentLevel = append([]int{node.value}, currentLevel...)
			} else {
				currentLevel = append(currentLevel, node.value)
			}

			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}

		}

		result = append(result, currentLevel)
		zigzag = !zigzag

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
