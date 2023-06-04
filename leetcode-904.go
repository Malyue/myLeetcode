package main

// 水果成篮

// 农场从左到右种植了一排果树，用整数数组fruits表示，其中fruits[i]是第i棵树上结的果子的种类
// 想要尽可能多地收集水果，然而，农场主人有以下规定：
/*
	只有 两个 篮子，并且每个篮子只能装 单一类型 的水果，每个篮子能装的水果没有限制
	可以选择任意一棵树开始采摘，必须从每棵树上恰好摘一个水果，采摘的水果应当符合篮子中的水果类型，每采摘依次，你将会向右移动到下一颗树，并继续采摘
	一旦走到某棵树前，但水果不符合篮子的水果类型，那么就必须停止采摘
*/

/*
	从样例比较容易看出如何解决

	输入：fruits =[1,2,1] 输出：3
	解释：可以采摘全部三棵树

	输入：fruits =[0,1,2,2] 输出：3
	解释：可以采摘[1,2,2]这三棵树，如果从第一棵树开始，那么只能采摘两棵树[0,1]
*/

func totalFruit(fruits []int) int {
	// 我们看样例，其实就是找最长的连续子数组，该子数组中只有两个不同的值

	// 用map来存储，key为水果的种类，value为该种类的数量
	// start指的是子数组的起始位置
	start := 0
	// 模拟篮子存储
	fruit_map := make(map[int]int)
	ans := 0
	// j指的是子数组的结束位置
	for j := 0; j < len(fruits); j++ {
		// 加入该种类
		fruit_map[fruits[j]]++
		if len(fruit_map) > 2 {
			// 如果长度超过2，那么就需要移动start指针了
			y := fruits[start]
			fruit_map[y]--
			if fruit_map[y] == 0 {
				delete(fruit_map, y)
			}
			start++
		}
		if j-start+1 > ans {
			ans = j - start + 1
		}
	}
	return ans
}
