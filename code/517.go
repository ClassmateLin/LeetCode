package main


var maxd, val int

func findBottomLeftValue(root *TreeNode) int {
	maxd, val = 0, 0
	dfs(root, 1)
	return val
}

// 深度优先搜索
func dfs(root *TreeNode, d int)  {
	if root == nil {
		return
	}
	if d > maxd {
		maxd = d
		val = root.Val
	}
	if root.Left != nil {
		dfs(root.Left, d+1)
	}
	if root.Right != nil {
		dfs(root.Right, d + 1)
	}
}

func main()  {
	
}