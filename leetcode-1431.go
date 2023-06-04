package main

// 拥有最多糖果的孩子
// 给你一个数组candies和一个整数extraCandies，其中candies[i]代表第i个孩子拥有的糖果数目
// 对每一个孩子，检查是否存在一种方案，将额外的extraCandies个糖果分配给孩子们之后，此孩子有最多的糖果
// 注意，允许有多个孩子同时拥有最多的糖果数目

// 样例
// 输入：candies = [2,3,5,1,3], extraCandies = 3
// 输出：[true,true,true,false,true]
// 解释：
// 第一个孩子在分配额外的3个糖果后，拥有了5个糖果，他/她有最多的糖果
// 第二个孩子在分配额外的3个糖果后，拥有了6个糖果，他/她有最多的糖果
// 第三个孩子在分配额外的3个糖果后，拥有了8个糖果，他/她有最多的糖果
// 第四个孩子在分配额外的3个糖果后，拥有了4个糖果，他/她没有最多的糖果
// 第五个孩子在分配额外的3个糖果后，拥有了5个糖果，他/她有最多的糖果
// 示例2
// 输入：candies = [4,2,1,1,2], extraCandies = 1
// 输出：[true,false,false,false,false]
// 解释：只有第一个孩子可以拥有最多的糖果
// 示例3
// 输入：candies = [12,1,12], extraCandies = 10
// 输出：[true,false,true]
// 提示
// 2 <= candies.length <= 100
// 1 <= candies[i] <= 100
// 1 <= extraCandies <= 50

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))

	// 最简单的解法，就是两遍遍历
	// 第一遍遍历找最大值，第二遍比较
	var max = -1
	for i := 0; i < len(candies); i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= max {
			result[i] = true
		}
	}

	return result
}
