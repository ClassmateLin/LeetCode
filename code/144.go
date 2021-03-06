package main

/*
给定一个二叉树，返回它的 前序 遍历。

 示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-preorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
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

		// 右子树先入栈，保证左子树先出栈
		stack = append(stack, s.Right)
		stack = append(stack, s.Left)
	}
	return res
}