package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if isLeaf(root.Left) { // 剪枝
		return root.Left.Val + sumOfLeftLeaves(root.Right);
	}
	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}


// 判断结点是否有子节点, 用于剪枝
func isLeaf(node *TreeNode) bool {
	if node == nil {
		return false
	}
	return node.Left == nil && node.Right == nil
}