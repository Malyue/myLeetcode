package main

import (
	"fmt"
	"strings"
)

// 反转字符串里的单词
// 给你一个字符串s，请你将s中的单词反转后输出
// 输入：s = "the sky is blue"
// 输出："blue is sky the"
// 输入：s = "  hello world  "
// 输出："world hello"

func reverseWords(s string) string {
	// 去掉头尾的空格
	s = strings.TrimSpace(s)
	// 中间可能有多个空格
	// 而Fields可以根据空格分割数组
	sarr := strings.Fields(s)
	result := ""
	for i := len(sarr) - 1; i >= 0; i-- {
		result = result + sarr[i] + " "
	}
	result = strings.TrimSpace(result)
	return result
}

func main() {
	fmt.Println(reverseWords("  Helood wriod   !  "))
}
