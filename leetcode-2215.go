package main

// 找出两数组的不同
// 给你两个下标从0开始的整数数组nums1和nums2，请你返回一个长度为2的列表answer
// 其中answer[0]是nums1中存在而nums2中不存在的数字
// answer[1]是nums2中存在而nums1中不存在的数字
// 示例1：
// 输入：nums1 = [1,2,3], nums2 = [2,4,6]
// 输出：[[1,3],[4,6]]
// 解释：
// 1和3存在于nums1中，但不存在于nums2中
// 4和6存在于nums2中，但不存在于nums1中

func findDifference(nums1 []int, nums2 []int) [][]int {
	// 用map来存储
	map1 := make(map[int]bool)
	map2 := make(map[int]bool)

	// 遍历存储到map中
	res := make([][]int, 2)
	for _, num := range nums1 {
		map1[num] = true
	}
	for _, num := range nums2 {
		map2[num] = true
	}

	for num := range map1 {
		if _, exist := map2[num]; !exist {
			res[0] = append(res[0], num)
		}
	}
	for num := range map2 {
		if _, exist := map1[num]; !exist {
			res[1] = append(res[1], num)
		}
	}
	return res
}
