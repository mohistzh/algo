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
Given a binary tree and a number sequence, find if the sequence is present as a root-to-leaf path in the given tree.
*/
func pathWithGivenSequence(root *MyTreeNode4, sequence []int) bool {
	var currentPath []int
	pathWithGivenSequenceRecursive(root, currentPath, sequence)
	return pathWithGivenSequenceRecursive2(root, sequence, 0)
}

// The solution when we reached to the leaf then compare sequence with path list
func pathWithGivenSequenceRecursive(currentNode *MyTreeNode4, currentPath []int, sequence []int) bool {
	result := false
	if currentNode == nil {
		return result
	}

	currentPath = append(currentPath, currentNode.value)
	if currentNode.left == nil && currentNode.right == nil {
		result = compareSlices(currentPath, sequence)
		currentPath = currentPath[:len(currentPath)-1]
		return result
	}
	return pathWithGivenSequenceRecursive(currentNode.left, currentPath, sequence) || pathWithGivenSequenceRecursive(currentNode.right, currentPath, sequence)

}

// comparing tree node with each element of sequence
func pathWithGivenSequenceRecursive2(currentNode *MyTreeNode4, sequence []int, sequenceIndex int) bool {
	if currentNode == nil {
		return false
	}
	if sequenceIndex >= len(sequence) || currentNode.value != sequence[sequenceIndex] {
		return false
	}

	if sequenceIndex == len(sequence)-1 && currentNode.left == nil && currentNode.right == nil {
		return true
	}

	return pathWithGivenSequenceRecursive2(currentNode.left, sequence, sequenceIndex+1) ||
		pathWithGivenSequenceRecursive2(currentNode.right, sequence, sequenceIndex+1)
}

func compareSlices(a []int, b []int) bool {
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	root := &MyTreeNode4{value: 1}
	root.left = &MyTreeNode4{value: 7}
	root.right = &MyTreeNode4{value: 9}
	root.right.left = &MyTreeNode4{value: 2}
	root.right.right = &MyTreeNode4{value: 9}
	fmt.Println(pathWithGivenSequence(root, []int{1, 9, 9}))
	root = &MyTreeNode4{value: 1}
	root.left = &MyTreeNode4{value: 0}
	root.right = &MyTreeNode4{value: 1}
	root.left.left = &MyTreeNode4{value: 1}
	root.right.left = &MyTreeNode4{value: 6}
	root.right.right = &MyTreeNode4{value: 5}
	fmt.Println(pathWithGivenSequence(root, []int{1, 0, 7}))
	fmt.Println(pathWithGivenSequence(root, []int{1, 1, 6}))
}
