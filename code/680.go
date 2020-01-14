package main

import "fmt"

/**
给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。

示例 1:

输入: "aba"
输出: True
示例 2:

输入: "abca"
输出: True
解释: 你可以删除c字符。
注意:

字符串只包含从 a-z 的小写字母。字符串的最大长度是50000。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-palindrome-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/**
- 可以使用左右指针遍历, 当左右指针分别指向的字符不相同时，删除一个字符再进行判断是否为回文串。
- 删除元素后判断回文串不需要判断整个字符串，之前的字符已经使用左右指针进行判断了，所以只要判断剩余的字符串是否回回文串即可。
- 删除字符串可以选择删除左指针指向的字符，也可以删除右指针指向的字符。
 */
func validPalindrome(s string) bool {
	i, j := 0, len(s) - 1
	for i < j  {
		if s[i] != s[j] {
			return isPalindrome(s, i, j - 1) || isPalindrome(s, i+1, j)
		}
		i++
		j--
	}
	return true
}

/**
左右指针遍历判断是否为回文串
 */
func isPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}


func main() {
	fmt.Println(validPalindrome("abscca"))
	fmt.Println(validPalindrome("abca"))
}
