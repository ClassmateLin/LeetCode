package main

import "math"

type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
}


func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right)))) + 1
}