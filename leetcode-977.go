package main

// 有序数组的平方

// 给定一个按非递减顺序排序的整数数组 nums，返回每个数字的平方组成的新数组，要求也按非递减顺序排序

func sortedSquares(nums []int) []int {
	// 由于是非递减顺序，所以可以知道，平方后最大的值一定是在两端
	// 所以可以使用双指针，分别指向两端，比较两端的值的平方，将较大的值放到新数组的最后
	// 然后将指针向中间移动，重复上述操作

	start, end := 0, len(nums)-1
	//创建一个新数组用来存储
	result := make([]int, len(nums))
	index := len(nums) - 1
	// 接下来要考虑的是，如果数组中有负数，那么平方后，最大的值可能在中间
	// 所以可以先找到数组中第一个非负数的位置，然后将start指向该位置，end指向该位置的下一个位置
	if len(nums) == 1 {
		result[index] = nums[0] * nums[0]
	}
	for {
		if start > end {
			break
		}
		if nums[start] >= 0 {
			// 如果大于0，则不需要再比较了,按序填充即可
			for i := end; i >= start; i-- {
				result[index] = nums[i] * nums[i]
				index--
			}
			break
		} else if nums[start]*nums[start] < nums[end]*nums[end] {
			result[index] = nums[end] * nums[end]
			end--
			index--
		} else {
			result[index] = nums[start] * nums[start]
			start++
			index--
		}
	}
	return result
}
