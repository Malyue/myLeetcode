package main

// 长度最小的子数组
// 给定一个含有n个正整数的数组和一个正整数target
// 找出数组中满足其和>=target的长度最小的连续子数组返回其长度，否则返回0

// 使用滑动窗口就是为了减小for循环，否则两个for循环也能解决问题
// 所以只用一个for循环，指向结尾，如果当前找到一个连续数组和超过target，那么这个起点可以直接跳过了，不用再往后找了
// 直接换一个起点，因为这个起点的值已经不可能是最小的了

func minSubArrayLen(target int, nums []int) int {
	// 滑动窗口
	// 本质是找到一个区间，使得区间内的和>=target
	// 如果大于该值，则左边减少，否则右边加入一个新值
	// 窗口起点
	i := 0
	// 窗口的值
	sum := 0
	result := len(nums) + 1
	// 窗口的终点位置
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		// 如果和超出了，就去掉最前面的值
		for sum >= target {
			subLength := j - i + 1
			if subLength < result {
				result = subLength
			}
			sum -= nums[i]
			i++
		}
	}
	if result == len(nums)+1 {
		return 0
	}
	return result
}
