package main

// 非递归实现二叉树的后序遍历
func postorderTraversal(root *TreeNode) []int {
	// 非递归实现二叉树的后序遍历
	res := []int{}

	stack :=[]*TreeNode{root} // 使用切片模拟栈

	for len(stack) > 0  {
		// 切片模拟出栈
		s := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		if s == nil {
		continue
		}
		res = append(res, s.Val)

		stack = append(stack, s.Left)// 左子树先入栈, 保证右子树先出站
		stack = append(stack, s.Right)
	}
	res = reverses(res)
	return res

}

// 切片翻转
func reverses(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}