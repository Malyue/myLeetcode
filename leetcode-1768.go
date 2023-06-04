package main

import (
	"fmt"
	"math"
)

// 交替合并字符串
// 给你两个字符串 word1 和 word2 。请你从 word1 开始，通过交替添加字母来合并字符串。
// 如果一个字符串比另一个字符串长，就将多出来的字母追加到合并后字符串的末尾。
// 返回 合并后的字符串 。
// 示例 1：
// 输入：word1 = "abc", word2 = "pqr"
// 输出："apbqcr"
// 解释：字符串合并情况如下所示：
// word1：  a   b   c
// word2：    p   q   r
// 合并后：  a p b q c r

func mergeAlternately(word1 string, word2 string) string {
	// 先获得较短的长度
	min := int(math.Min(float64(len(word1)), float64(len(word2))))
	// 在min长度以内交替，后续直接追加
	var result string
	for i := 0; i < min; i++ {
		result += string(word1[i])
		result += string(word2[i])
	}
	if len(word1) == min {
		// 追加word2
		result += word2[min:]
	} else {
		result += word1[min:]
	}
	return result
}

func main() {
	fmt.Println(mergeAlternately("abc", "pqrd"))
}
