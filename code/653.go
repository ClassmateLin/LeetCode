package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 在二叉树中查找两个节点使之和等于k
func findTarget(root *TreeNode, k int) bool {
	nums := []int{}
	toArray(root, &nums)
	i, j := 0, len(nums) - 1
	for i < j {
		sum := nums[i] + nums[j]
		if sum == k {
			return true
		}
		if sum < k {
			i++
		}else {
			j--
		}
	}
	return false
}


// 二叉查找树中序遍历后得到一个有序数组
func toArray(root *TreeNode, arr *[]int){
	if root == nil {
		return
	}
	toArray(root.Left, arr)
	*arr = append(*arr, root.Val)
	toArray(root.Right, arr)
}

func main()  {
	root := TreeNode{Val:2}
	root.Left = &TreeNode{Val:1}
	root.Right = &TreeNode{Val:3}
	findTarget(&root, 1)
}