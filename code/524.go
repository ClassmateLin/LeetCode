package main

import (
	"fmt"
	"strings"
)

/**
给定一个字符串和一个字符串字典，找到字典里面最长的字符串，该字符串可以通过删除给定字符串的某些字符来得到。如果答案不止一个，返回长度最长且字典顺序最小的字符串。如果答案不存在，则返回空字符串。

示例 1:

输入:
s = "abpcplea", d = ["ale","apple","monkey","plea"]

输出:
"apple"
示例 2:

输入:
s = "abpcplea", d = ["a","b","c"]

输出:
"a"
说明:

所有输入的字符串只包含小写字母。
字典的大小不会超过 1000。
所有输入的字符串长度不会超过 1000。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-word-in-dictionary-through-deleting
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


func findLongestWord(s string, d []string) string {
	var longestWord string
	for _, target := range d {
		l1, l2 := len(longestWord), len(target)
		// 保存的最长序列和当前序列的长度，　以及判断字典序大小
		if l1 > l2 || (l1 == l2 && strings.Compare(longestWord, target) == -1){
			continue
		}
		// 如果是子序列则更新
		if isSubstr(s, target) {
			longestWord = target
		}
	}
	return longestWord
}


/**
使用双指针判断某字符串是否为另一字符串的子序列
 */
func isSubstr(s string, target string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(target) {
		if s[i] == target[j] {
			j++
		}
		i++
	}
	return j == len(target)
}

func main() {
	s := "abpcplea"
	d := []string{"ale","apple","monkey","plea"}
	fmt.Println(findLongestWord(s, d))
}
