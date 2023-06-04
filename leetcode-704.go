package main

// 二分查找

func search(nums []int, target int) int {
	// 给定的是有序数组 `nums`
	// 要找到给定的target对应的下标
	// 如果找不到则返回-1
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return 1
		}
	}

	// 二分查找，主要要学会判断区间
	// 一开始区间是[0,len(nums)-1]
	// 且每次都是取中间的值，然后判断中间的值和target的大小，即下次选择的也还是mid-1,因为mid已经判断完了
	left, right := 0, len(nums)-1
	for {
		if left > right {
			return -1
		}
		// 这里要避免溢出，所以不能使用(left+right)/2
		// 这里要使用位运算,因为left+(right-left)/2也有溢出的可能,且移位运算会快
		middle := left + ((right - left) >> 1)
		if nums[middle] == target {
			return middle
		}

		if nums[middle] > target {
			// 往左边找
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	println(search(nums, target))
}
