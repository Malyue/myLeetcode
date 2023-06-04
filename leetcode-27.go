package main

import "fmt"

// 删除数组中的指定元素
// 给定一个数组 nums 和 一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度

func removeElement(nums []int, val int) int {
	// 双指针
	// 本质是将不等于val的移动到数组左边
	// 因为最后需要返回数组的长度，所以可以将不等于val的值移动到数组左边，等于val的值移动到数组右边
	p, q := 0, len(nums)-1
	for {
		if p > q {
			break
		}
		if nums[p] == val {
			nums[p] = nums[q]
			q--
		} else {
			p++
		}
	}
	return len(nums[:p])
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	val := 3
	fmt.Println(removeElement(nums, val))
}
