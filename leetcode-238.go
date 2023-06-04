package main

// 除本身以外数组的乘积
// 给你一个整数数组nums，返回数组answer，其实answer[i]等于nums中除nums[i]之外其余各元素的乘积
// 题目数据保证数组 nums 之中任意元素的全部前缀元素和后缀的乘积都在32位整数范围内
// 请不要使用除法，且在O(n)时间复杂度内完成此题

func productExceptSelf(nums []int) []int {
	// 开始最简单的想法，就是每个都遍历一遍，但是这样需要O(n^len(nums))
	// 不然就是所有都相乘，然后每次除去对应的那个值，这样也能解决，而且只需要O(n),但是不允许除法
	// 那就考虑用map来保存，或者用两个数组来保存某个位置的前缀和后缀各自的乘积
	// 时间复杂度O(n),空间复杂度O(n)
	//length := len(nums)
	//L, R, answer := make([]int, length), make([]int, length), make([]int, length)
	//// 有两个特殊情况，nums[0]没有前缀，nums[length-1]没有后缀
	//L[0] = 1
	//R[length-1] = 1
	//for i := 1; i < length; i++ {
	//	L[i] = L[0] * nums[i-1]
	//}
	//
	//for i := length - 2; i >= 0; i-- {
	//	R[i] = R[i+1] * nums[i+1]
	//}
	//
	//for i := 0; i < length; i++ {
	//	answer[i] = L[i] * R[i]
	//}
	//
	//return answer

	// 时间复杂度O(n),空间复杂度O(1)
	length := len(nums)
	answer := make([]int, length)
	answer[0] = 1
	for i := 1; i < length; i++ {
		answer[i] = nums[i-1] * answer[i-1]
	}
	R := 1
	for i := length - 1; i >= 0; i-- {
		answer[i] = answer[i] * R
		R *= nums[i]
	}

	return answer
}
