package main

// 给定字符串s和t，判断s是否为t的子序列
// 字符串的一个子序列是原始字符串删除一些(也可以不删除)字符而不改变剩余字符相对位置形成的新字符串。
// 例如，"ace"是"abcde"的一个子序列，而"aec"不是
// 进阶：如果有大量输入的S，称作S1,S2,...,Sk其中k>=10亿，你需要依次检查它们是否为T的子序列
// 在这种情况下，你会怎样改变代码？
// 示例1：
// 输入：s = "abc", t = "ahbgdc"
// 输出：true
// 示例2：
// 输入：s = "axc", t = "ahbgdc"
// 输出：false

func isSubsequence(s string, t string) bool {
	//// 最简单的，遍历原始字符串，然后子串一个指针遍历，当有一个对上了，子串指针往后走
	//if len(s) == 0 {
	//	return true
	//}
	//if len(t) == 0 {
	//	return false
	//}
	//sub := 0
	//for i := 0; i < len(t); i++ {
	//	// 如果相等则sub+1
	//	if s[sub] == t[i] {
	//		sub++
	//		if sub == len(s) {
	//			return true
	//		}
	//	}
	//}
	//return false

	// 动态规划，前面一种方法，我们很明显知道较多的时间是在找下一个匹配字符
	// 令f[i][j]表示字符串t从位置i开始往后字符j第一次出现的位置，在进行状态转移时，如果t中位置i的字符就是j，那么f[i][j] = i,
	// 否则j出现在位置i+1开始往后，即f[i][j] = f[i+1][j],因此我们要倒过来进行动态规划，从后往前枚举i
	// 状态转移方程:
	// f[i][j] = i,t[i]=j
	//  	   = f[i+1][j],t[i] != j
	// 假定下标从0开始，那么f[i][j] 中有0<=i<=m-1
	// 其实就是KMP算法，明天一定学！！！

	n, m := len(s), len(t)
	f := make([][26]int, m+1)
	// 表示
	for i := 0; i < 26; i++ {
		f[m][i] = m
	}
	for i := m - 1; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			if t[i] == byte(j+'a') {
				f[i][j] = i
			} else {
				f[i][j] = f[i+1][j]
			}
		}
	}
	add := 0
	for i := 0; i < n; i++ {
		if f[add][int(s[i]-'a')] == m {
			return false
		}
		add = f[add][int(s[i]-'a')] + 1
	}
	return true
}
