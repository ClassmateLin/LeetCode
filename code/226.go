package main


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// go语言特性交换值
	root.Left, root.Right = root.Right, root.Left

	invertTree(root.Left)
	invertTree(root.Right)
	return root
}