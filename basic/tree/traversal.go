package tree

import "fmt"

// 前序遍历
func (n *TreeNode) PreOrderTraversal() {
	if n == nil {
		return
	}
	fmt.Printf("%d ", n.Val)
	n.Left.PreOrderTraversal()
	n.Right.PreOrderTraversal()
}

// 中序遍历
func (n *TreeNode) InOrderTraversal() {
	if n == nil {
		return
	}
	n.Left.InOrderTraversal()
	fmt.Printf("%d ", n.Val)
	n.Right.InOrderTraversal()
}

// 后序遍历
func (n *TreeNode) PostOrderTraversal() {
	if n == nil {
		return
	}
	n.Left.PostOrderTraversal()
	n.Right.PostOrderTraversal()
	fmt.Printf("%d ", n.Val)
}

// 层序遍历
func (n *TreeNode) LevelOrderTraversal() {
	if n == nil {
		return
	}

	queue := []*TreeNode{n}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}
