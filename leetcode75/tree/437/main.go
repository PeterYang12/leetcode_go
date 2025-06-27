package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	// 前缀和
	preSum := map[int64]int{0: 1}
	ans := 0
	var dfs func(root *TreeNode, curr int64)
	dfs = func(root *TreeNode, curr int64) {
		if root == nil {
			return
		}
		curr += int64(root.Val)
		ans += preSum[curr-int64(targetSum)]
		preSum[curr] += 1
		dfs(root.Left, curr)
		dfs(root.Right, curr)
		preSum[curr] -= 1
	}
	dfs(root, 0)
	return ans
}
func main() {

}
