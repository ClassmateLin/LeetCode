package main


// 判断是否有子树
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	return isSame(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}


// 判断两树是否相同
func isSame(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	return s.Val == t.Val && isSame(s.Left, t.Left) && isSame(s.Right, t.Right)
}