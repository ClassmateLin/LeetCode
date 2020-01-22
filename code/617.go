package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 合并二叉树
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	root := new(TreeNode)
	root.Val = t1.Val + t2.Val
	root.Left = mergeTrees(t1.Left, t2.Left) // 合并左子树
	root.Right = mergeTrees(t1.Right, t2.Right) // 合并右子树

	return root
}
