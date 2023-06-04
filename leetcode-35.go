package main

func searchInsert(nums []int, target int) int {
	// 给定一个排序数组和一个目标，在数组中找到目标，并返回其索引。
	// 如果未找到，则返回按顺序会插入的位置
	// 二分查找，主要要学会判断区间
	// 一开始区间是[0,len(nums)-1]
	// 且每次都是取中间的值，然后判断中间的值和target的大小，即下次选择的也还是mid-1,因为mid已经判断完了
	// 其实归为一句话，这道题就是寻找第一个不小于 target 的值的位置
	left, right := 0, len(nums)-1
	if nums[right] < target {
		return right + 1
	}
	for {
		// 收缩区间
		// 比如区间小于2，那么就可以直接判断返回了
		if left < right {
			return left
		}
		mid := left + ((right - left) >> 1)
		// 寻找第一个不小于target的位置
		// 区间是[left,right],要找的值由于是不小于，所以要包含mid
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 2
	println(searchInsert(nums, target))
}
