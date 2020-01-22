package main

// 判断是否为镜像树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSmt(root.Left, root.Right)
}


// 判断是否为镜像树
func isSmt(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil ||  t2 == nil  {
		return false
	}

	// t1的左子树跟t2的右子树进行对比
	return  (t1.Val == t2.Val) && isSmt(t1.Left, t2.Right) && isSmt(t1.Right, t2.Left)
}