package main

// 判断路径和是否等于一个数
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && root.Val == sum {
		return true
	}
	return hasPathSum(root.Left, sum - root.Val) || hasPathSum(root.Right, sum - root.Val)
}