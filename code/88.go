package main

/*
给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。

说明:

初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
示例:

输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

输出: [1,2,2,3,5,6]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/**
- 设两指针分别指向nums1和nums2，对分别指向的值进行判断。
- 题目要求将数组合并到nums1中，所以要从数组末尾开始遍历，否则会将nums1的数据覆盖掉。
 */
func merge(nums1 []int, m int, nums2 []int, n int)  {
	p := m + n -1
	m, n = m - 1, n - 1  // m, n 是数组长度, 减一当索引用
	for m>=0 || n>=0{
		if m < 0 { // nums1已经合并完了, 剩余的都是nums2数据
			nums1[p] = nums2[n]
			n--
		} else if n < 0 { // nums2已经合并完了, 剩余的都是nums1数组
			nums1[p] = nums1[m]
			m--
		} else if nums1[m] > nums2[n] { //　大的数先合并
			nums1[p] = nums1[m]
			m--
		} else {
			nums1[p] = nums2[n]
			n--
		}
		p--  // 移至下一次需合并的位置
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m := 3
	n := 3
	merge(nums1, m, nums2, n)
}
