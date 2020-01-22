package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil { //根节点为空，找不到
		return -1
	}
	// 左右子树为空找不到
	if root.Left == nil && root.Right == nil {
		return -1
	}
	left := root.Left.Val
	right := root.Right.Val
	if left == root.Val { // 左子树与根节点值相同继续往下找
		left = findSecondMinimumValue(root.Left)
	}
	if right == root.Val { // 右子树与根节点值相同继续往下找
		right = findSecondMinimumValue(root.Right)
	}
	// 左右子树进行比较，返回值小的
	if left != -1 && right != -1 {
		return min(left, right)
	}
	// 左子树不为空，则返回左子树，否则为右子树
	if left != -1 {
		return left
	}
	return right
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}