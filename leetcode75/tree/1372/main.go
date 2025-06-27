package main

import (
	"fmt"
	"leetcode_go/basic/tree"
)

type TreeNode = tree.TreeNode

func longestZigZag1(root *TreeNode) int {
	maxLength := 0
	// 0, left, 1, right
	var dfs func(root *TreeNode, dir int) int
	dfs = func(root *TreeNode, dir int) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left, 0)
		right := dfs(root.Right, 1)
		currentMax := max(left, right)
		maxLength = max(maxLength, currentMax)
		if dir == 0 {
			return right + 1
		} else {
			return left + 1
		}
	}
	dfs(root, -1)
	return maxLength
}
func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	var dfs func(root *TreeNode, dir, length int)
	dfs = func(root *TreeNode, dir, length int) {
		if root == nil {
			return
		}
		ans = max(ans, length)
		if dir == 0 {
			dfs(root.Right, 1, length+1)
			dfs(root.Left, 0, 1)
		} else {
			dfs(root.Left, 0, length+1)
			dfs(root.Right, 1, 1)
		}
	}
	dfs(root.Left, 0, 1)
	dfs(root.Right, 1, 1)
	return ans
}
func main() {
	arr := []interface{}{1, nil, 1, 1, 1, nil, nil, 1, 1, nil, 1, nil, nil, nil, 1}
	//arr := []interface{}{1, 1, 1, nil, 1, nil, nil, 1, 1, nil, 1}
	root := tree.BuildTreeFromArray(arr)
	fmt.Println(longestZigZag(root))
}
