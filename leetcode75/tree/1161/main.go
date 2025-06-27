package main

import (
	"fmt"
	"leetcode_go/basic/tree"
	"math"
)

type TreeNode = tree.TreeNode

func maxLevelSum(root *TreeNode) int {
	ans := 0 //至少有一个节点
	maxSum := math.MinInt
	queue := []*TreeNode{root}
	index := 0
	for len(queue) > 0 {
		index++
		curSum := 0
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[i]
			curSum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[n:]
		if curSum > maxSum {
			maxSum = curSum
			ans = index
		}
	}
	return ans
}
func main() {
	arr := []interface{}{1, 7, 0, 7, -8, nil, nil}
	root := tree.BuildTreeFromArray(arr)
	fmt.Println(maxLevelSum(root))
}
