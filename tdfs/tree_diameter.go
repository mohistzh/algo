package main

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
func findTreeDiameter(root *MyTreeNode6) int {
	return -1
}

func main() {
	root := &MyTreeNode6{value: 1}
	root.left = &MyTreeNode6{value: 2}
	root.right = &MyTreeNode6{value: 3}
	root.right.left = &MyTreeNode6{value: 5}
	root.right.right = &MyTreeNode6{value: 6}
	root.right.right = &MyTreeNode6{value: 9}
	root.right.right = &MyTreeNode6{value: 11}
	root.right.left.left = &MyTreeNode6{value: 7}
	root.right.left.right = &MyTreeNode6{value: 8}
	root.right.left.right.right = &MyTreeNode6{value: 10}

}
