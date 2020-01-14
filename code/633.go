package main

import (
	"fmt"
	"math"
)
/**

633. 平方数之和
给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a2 + b2 = c。

示例1:

输入: 5
输出: True
解释: 1 * 1 + 2 * 2 = 5
 

示例2:

输入: 3
输出: False

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sum-of-square-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

 */

/**
- 两数之和与两数字平方和类似，可以使用双指针遍历。
- 设左指针变量为i, 右指针变量为j。i初始为0, 使遍历次数最少，j 应该为 target的开平方。
- 左右指针的平方和为: powSum = i * i + j * j。
- 如果 powSum < target则左指针右移。
- 如果 powSum > target则右指针左移。
- 如果 powSum = target　则返回True。
- 其他返回Flase。
 */
func judgeSquareSum(c int) bool {
	if c < 0 {
		return false
	}
	i, j := 0, int(math.Sqrt(float64(c)))

	for i <= j {
		powSum := i * i + j * j
		if powSum == c {
			return true
		}else if powSum < c {
			i ++
		}else {
			j--
		}
	}
	return false
}

func main() {
	fmt.Println(judgeSquareSum(10))
}
