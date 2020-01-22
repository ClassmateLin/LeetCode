package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 求二叉树的高度
func Height(node *TreeNode) (h int)  {
	if node == nil {
		return 0
	}

	left := Height(node.Left)
	right := Height(node.Right)

	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

// 判断是否为二叉平衡树
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	diff := Height(root.Left) - Height(root.Right)

	if diff < -1 || diff > 1 {
		return false
	}

	if isBalanced(root.Left) && isBalanced(root.Right) {
		return true
	}
	return false
}