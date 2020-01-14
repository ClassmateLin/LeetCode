package main

import "fmt"

/**
编写一个函数，以字符串作为输入，反转该字符串中的元音字母。

示例 1:

输入: "hello"
输出: "holle"
示例 2:

输入: "leetcode"
输出: "leotcede"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-vowels-of-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func reverseVowels(s string) string {
	r := []rune(s) // 字符串不可修改，转为rune类型
	// 存储所有的元音字母
	vowelLetter := map[rune]int{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1, 'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U':1}
	i, j := 0, len(s) - 1
	for i <= j {
		if vowelLetter[r[i]] != 1{ // 如果左指针指向的非元音字符，　则右移
			i++
		}else if vowelLetter[r[j]] != 1 { // 如果右指针指向的为非元音字符，则左移
			j--
		}else { // 元音字母进行交换
			r[i], r[j] = r[j], r[i]
			i++
			j--
		}
	}
	return string(r)
}

func main() {
	fmt.Println(reverseVowels("LeetCode"))
}
