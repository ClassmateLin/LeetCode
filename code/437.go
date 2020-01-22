package main


func pathSum(root *TreeNode, sum int) int {
	if root == nil{
		return 0
	}

	return pathSumRecursion(root, sum, 0) + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func pathSumRecursion(root *TreeNode, sum int, count int)int{
	if root == nil {
		return count
	}

	sum -= root.Val
	if sum == 0{
		count++
	}
	count = pathSumRecursion(root.Left, sum, count)
	count = pathSumRecursion(root.Right, sum, count)

	return count
}
