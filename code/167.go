package main

import "fmt"

/**
两数之和 II - 输入有序数组

给定一个已按照升序排列 的有序数组，找到两个数使得它们相加之和等于目标数。

函数应该返回这两个下标值 index1 和 index2，其中 index1 必须小于 index2。

说明:

返回的下标值（index1 和 index2）不是从零开始的。
你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。
示例:

输入: numbers = [2, 7, 11, 15], target = 9
输出: [1,2]
解释: 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/**
- 数组nums为有序数组, 可以使用左右指针，一个初始指向最小值(左指针)，一个初始指向最大值(右指针)，
左指针往右移动，右指针往左移动。
- 设左指针变量为i, 右指针变量为j, 目标值为target。
- 如果`nums[i] + nums[j] == target` 则得到所求解。
- 如果`nums[i] + nums[j] > target`，那么移动右指针。
- 如果`nums[i] + nums[j] < target`，那么移动左指针。
- 不满足以上三种情况则找不到结果。
 */
func twoSum(numbers []int, target int) []int {
	res := make([]int, 2)
	i, j := 0, len(numbers) - 1
	for i < j  {
		if numbers[i] + numbers[j] == target {
			res[0], res[1] = i + 1, j + 1 // 注意题目中索引是从１开始的
			return res
		}else if numbers[i] + numbers[j] < target {
			i++
		}else {
			j--
		}
	}
	return res
}


func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
