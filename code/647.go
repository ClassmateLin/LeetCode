package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	res := []float64{}
	if root == nil {
		return res
	}
	cnt, sum := 0, 0 // 统计节点数和节点值和
	q := []*TreeNode{root} // 切片当队列使用

	for len(q) > 0 {
		cnt = len(q)
		sum = 0
		for i := 0; i < cnt; i++ { // 遍历当前层的节点，并将下层节点添加到队列
			tmp := q[0]
			q = q[1:]
			if tmp.Left != nil { // 左子树入队
				q = append(q, tmp.Left)
			}
			if tmp.Right != nil { // 右子树入队
				q = append(q, tmp.Right)
			}
			sum += tmp.Val
		}
		// 计算该层的平均值
		res = append(res, float64(sum) / float64(cnt))
	}
	return res
}