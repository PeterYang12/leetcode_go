package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var list []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			list = append(list, node.Val)
			return
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root1)
	list1 := append([]int{}, list...)
	list = []int{}
	dfs(root2)
	if len(list1) != len(list) {
		return false
	}
	for i := 0; i < len(list1); i++ {
		if list1[i] != list[i] {
			return false
		}
	}
	return true
}

func main() {

}
