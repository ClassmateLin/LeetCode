package main

import (
	"fmt"
)

/**
编写一个函数，输入是一个无符号整数，返回其二进制表达式中数字位数为 ‘1’ 的个数（也被称为汉明重量）。
示例 1：
输入：00000000000000000000000000001011
输出：3
解释：输入的二进制串 00000000000000000000000000001011 中，共有三位为 '1'。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-1-bits
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func hammingWeight(num uint32) int {
	ans := 0
	for num != 0 {
		if num & 1 == 1{ //让num与1进行按位与运算，取得num最低位判断是否位1
			ans += 1
		}
		num >>= 1
	}
	return ans
}

func main()  {
	var a uint32 = 10
	fmt.Println(hammingWeight(a))
	a = 11
	fmt.Println(hammingWeight(a))
	a = 12
	fmt.Println(hammingWeight(a))
}