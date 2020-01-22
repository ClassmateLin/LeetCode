package main

// 求二叉搜索树第k小的元素
func kthSmallest(root *TreeNode, k int) int {
	leftCnt := count(root.Left)
	if leftCnt == k - 1 {
		return root.Val
	}
	if leftCnt > k - 1 {
		return kthSmallest(root.Left, k)
	}
	return kthSmallest(root.Right, k - leftCnt - 1)
}


func count(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + count(node.Left) + count(node.Right)
}