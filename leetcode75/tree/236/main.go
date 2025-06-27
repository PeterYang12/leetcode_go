package main

//import t "leetcode_go/basic/tree"
import "leetcode_go/basic/tree"

type TreeNode = tree.TreeNode

// type TreeNode tree.TreeNode  // 新类型定义
//
//	func main() {
//		var n1 tree.TreeNode
//		var n2 TreeNode
//		n1 = tree.TreeNode(n2)  // 必须显式转换
//		n2 = TreeNode(n1)       // 反向同理
//	}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}
func main() {

}
