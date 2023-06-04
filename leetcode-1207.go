package main

// 独一无二的出现次数
// 给你一个整数数组arr，请你帮忙统计数组中每个数的出现次数
// 如果每个数的出现次数都是独一无二的，就返回true；否则返回false
// 示例1：
// 输入：arr = [1,2,2,1,1,3]
// 输出：true
// 解释：在该数组中，1出现了3次，2出现了2次，3只出现了1次，没有两个数的出现次数相同
// 示例2：
// 输入：arr = [1,2]
// 输出：false
// 示例3：
// 输入：arr = [-3,0,1,-3,1,1,1,-3,10,0]
// 输出：true

func uniqueOccurrences(arr []int) bool {
	res := make(map[int]int)
	for _, value := range arr {
		if _, exist := res[value]; exist {
			res[value] = res[value] + 1
		} else {
			res[value] = 1
		}
	}
	// 新建一个hash存储出现数量
	mapNum := make(map[int]bool)
	for _, value := range res {
		if _, exist := mapNum[value]; exist {
			return false
		}
		mapNum[value] = true
	}
	return true
}

func main() {
	uniqueOccurrences([]int{1, 2})
}
