package main


// 非递归实现二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	res := []int{}

	if root == nil {
		return res
	}

	stack := []*TreeNode{}
	cur := root

	for cur != nil || len(stack) > 0 {
		for cur != nil  { // 先将左子树进栈
			stack = append(stack, cur)
			cur = cur.Left
		}
		// 取栈顶元素，即为根
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)

		// 遍历右子树
		cur = node.Right
	}
	return res
}