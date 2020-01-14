package main

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	arr := []int{}
	str := ""
	i := len(num1) - 1
	j := len(num2) - 1
	cur := 0
	for i>=0 || j>=0 || cur != 0 {
		if(i>=0){
			// 题目不能直接将num2转int， 这里c++/c可以直接运算，但是go不行
			cur += int(num1[i] - '0')
			i--
		}
		if(j>=0){
			cur += int(num2[j] - '0')
			j--
		}
		arr = append(arr, cur%10)
		cur /= 10
	}

	for k:=len(arr)-1; k>=0; k--{
		str += strconv.Itoa(arr[k])
	}
	return str
}

func main() {
	fmt.Println(addStrings("11", "9"))
}
