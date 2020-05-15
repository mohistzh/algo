package main

// MyTreeNode9 tree node
type MyTreeNode9 struct {
	value             int
	left, right, next *MyTreeNode9
}

/*
Given a binary tree, connect each node with its level order successor.
The last node of each level should point to the first node of the next level.

*/
func connectAllLevelOrderSiblings(root *MyTreeNode9) {
	var queue []*MyTreeNode9
	queue = append(queue, root)
	var currentNode *MyTreeNode9
	var previousNode *MyTreeNode9
	for len(queue) > 0 {
		currentNode = queue[0]
		queue = queue[1:]
		if previousNode != nil {
			previousNode.next = currentNode
		}
		previousNode = currentNode
		if currentNode.left != nil {
			queue = append(queue, currentNode.left)
		}
		if currentNode.right != nil {
			queue = append(queue, currentNode.right)
		}

	}
}

func printMyTreeNode9(root *MyTreeNode9) {

}

func main() {
	root := &MyTreeNode9{value: 1}
	root.left = &MyTreeNode9{value: 2}
	root.right = &MyTreeNode9{value: 3}
	root.left.left = &MyTreeNode9{value: 4}
	root.left.right = &MyTreeNode9{value: 5}
	root.right.left = &MyTreeNode9{value: 6}
	root.right.right = &MyTreeNode9{value: 7}
	connectAllLevelOrderSiblings(root)
	printMyTreeNode9(root)
}
