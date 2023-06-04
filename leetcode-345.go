package main

import "strings"

// 反转字符串中的元音字母
// 编写一个函数，以字符串作为输入，反转该字符串中的元音字母
// 元音字母包括a、e、i、o、u，但不包括y
// 示例1：
// 输入："hello"
// 输出："holle"
// 示例2：
// 输入："leetcode"
// 输出："leotcede"
// 提示：
// 元音字母不包含字母"y"

// 反转其实有点没看懂，看了一下示例简单先理解为从前往后的第n个和从后往前的第n个元音字母进行交换吧
func reverseVowels(s string) string {
	// 双指针
	t := []byte(s)
	n := len(s)
	i, j := 0, n-1
	for i < j {
		for i < n && !strings.Contains("aeiouAEIOU", string(t[i])) {
			i++
		}
		for j > 0 && !strings.Contains("aeiouAEIOU", string(t[j])) {
			j--
		}
		if i < j {
			t[i], t[j] = t[j], t[i]
			i++
			j--
		}
	}
	return string(t)
}

func IfYuanyin(s string) bool {
	if s == "A" || s == "E" || s == "I" || s == "O" || s == "U" || s == "a" || s == "e" || s == "i" || s == "o" || s == "u" {
		return true
	}
	return false
}
