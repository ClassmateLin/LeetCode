package main

import "fmt"

/*
编写一个算法来判断一个数是不是“快乐数”。

一个“快乐数”定义为：对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和，
然后重复这个过程直到这个数变为 1，也可能是无限循环但始终变不到 1。
如果可以变为 1，那么这个数就是快乐数。

示例: 

输入: 19
输出: true
解释:
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/happy-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/*
出现无限循环说明有重复数字,
算出新的数如果不等于1且不重复则添加到map,
否则返回true/false
 */
func isHappy(n int) bool {
	m := map[int]int{}
	for n != 1 {
		n = calNum(n)
		val, ok := m[n]
		if ok && val>0{
			return false
		}else {
			m[n] = 1
		}
	}
	return true
}

/*
计算新的数
 */
func calNum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n%10) * (n%10)
		n /= 10
	}
	return sum
}


func main() {
	r1 := isHappy(19)
	r2 := isHappy(21)
	fmt.Println(r1)
	fmt.Println(r2)
}
