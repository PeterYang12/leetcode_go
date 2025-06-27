package main

import (
	"fmt"
	"leetcode_go/basic/tree"
)

type TreeNode = tree.TreeNode

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{root}
	var ans []int
	for len(queue) > 0 {
		n := len(queue)
		var tmpQueue []*TreeNode
		for i := 0; i < n; i++ {
			node := queue[i]
			if node.Left != nil {
				tmpQueue = append(tmpQueue, node.Left)
			}
			if node.Right != nil {
				tmpQueue = append(tmpQueue, node.Right)
			}

			if i == n-1 {
				ans = append(ans, node.Val)
			}
		}
		queue = tmpQueue
	}
	return ans
}
func main() {
	arr := []interface{}{1, 2, 3, nil, 5, nil, 4}
	root := tree.BuildTreeFromArray(arr)
	fmt.Println(rightSideView(root))
}
