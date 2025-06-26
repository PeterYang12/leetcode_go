package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	ans := 0
	var dfs func(node *TreeNode, maxPath int)
	dfs = func(node *TreeNode, maxPath int) {
		if node == nil {
			return
		}

		if node.Val >= maxPath {
			ans++
		}

		dfs(node.Left, max(maxPath, node.Val))
		dfs(node.Right, max(maxPath, node.Val))
	}
	dfs(root, math.MinInt32)
	return ans
}

func main() {

}
