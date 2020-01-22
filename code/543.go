package main

import "math"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var max int = 0

func diameterOfBinaryTree(root *TreeNode) int {
	max = 0 // 运行样例前需要重置为0
	Depth(root)
	return max
}

// 最长路径
func Depth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := Depth(node.Left)
	right := Depth(node.Right)

	if left + right > max {
		max = left + right
	}

	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}
