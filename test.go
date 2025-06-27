package main

import (
	"basic/basic/tree"
	"fmt"
)

func main() {
	// 使用数组构造二叉树
	// nil 表示空节点
	arr := []interface{}{1, 2, 3, 4, nil, 5, 6, nil, nil, 7}
	root := tree.BuildFromArray(arr)

	fmt.Println("层序遍历结果(二维数组，每层一个子数组):")
	fmt.Println(root.LevelOrder())

	fmt.Println("\n前序遍历结果:")
	fmt.Println(root.PreOrder())

	fmt.Println("\n中序遍历结果:")
	fmt.Println(root.InOrder())

	fmt.Println("\n后序遍历结果:")
	fmt.Println(root.PostOrder())

	fmt.Println("\n字符串表示(层序):")
	fmt.Println(root)
}
