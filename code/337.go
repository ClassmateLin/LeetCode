package main


// 间隔遍历
func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	val1 := root.Val

	// 打劫子孙辈节点
	if root.Left != nil {
		val1 += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		val1 += rob(root.Right.Left) + rob(root.Right.Right)
	}

	// 打劫左右子树
	val2 := rob(root.Left) + rob(root.Right)

	if val1 > val2 {
		 return val1
	}
	return val2
}


func main()  {

}