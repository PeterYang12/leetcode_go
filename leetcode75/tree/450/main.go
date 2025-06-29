package main

import (
	"leetcode_go/basic/tree"
)

type TreeNode = tree.TreeNode

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else if root.Left == nil || root.Right == nil {
		// root.Val == key
		if root.Left == nil {
			root = root.Right
		} else {
			root = root.Left
		}
	} else {
		successor := root.Right
		// find the smallest one
		for successor.Left != nil {
			successor = successor.Left
		}
		successor.Right = deleteNode(root.Right, successor.Val)
		successor.Left = root.Left
		return successor
	}
	return root
}
func main() {

}
