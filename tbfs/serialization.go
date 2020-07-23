package main

import (
	"bytes"
	"fmt"
	"strconv"
)

//MyTreeNode11 simple tree struct
type MyTreeNode11 struct {
	value       int
	left, right *MyTreeNode11
}

// preorderSerialize serialize the given tree node to a string by PreOrder traversal
func preorderSerialize(root *MyTreeNode11) string {
	var buf bytes.Buffer
	_preorderSerialize(root, &buf)
	return buf.String()
}

// preorderSerialize serialize the given tree node to a string by PreOrder traversal
func _preorderSerialize(root *MyTreeNode11, buf *bytes.Buffer) {
	if root == nil {
		buf.WriteString("#")
		buf.WriteString(",")
		return
	}
	buf.WriteString(strconv.Itoa(root.value))
	buf.WriteString(",")
	_preorderSerialize(root.left, buf)
	_preorderSerialize(root.right, buf)
}

// preorderDeserialize deserialize the given string to a Tree
func preorderDeserialize(serialized string) *MyTreeNode11 {
	return nil
}

func main() {
	root := &MyTreeNode11{value: 1}
	root.left = &MyTreeNode11{value: 2}
	root.right = &MyTreeNode11{value: 3}
	root.left.right = &MyTreeNode11{value: 4}
	serializedStr := preorderSerialize(root)
	fmt.Println(serializedStr)
}
