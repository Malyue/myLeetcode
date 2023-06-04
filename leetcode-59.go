package main

// 给定一个正整数n，生成一个包含 1 到 n² 所有元素，且元素按顺时针顺序螺旋排列n*n的正方形矩阵matrix

func generateMatrix(n int) [][]int {
	// 如果按照正常一圈一圈赋值，不知道要实现多久，有很多情况要处理的
	matrix := make([][]int, n)
	startx, starty := 0, 0
	var loop int = n / 2
	var center int = n / 2
	count := 1
	offset := 1
	// 初始化每一行
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	for loop > 0 {
		i, j := startx, starty

		// 行数不变，列数在变
		for j = starty; j < n-offset; j++ {
			matrix[startx][j] = count
			count++
		}

		// 列数不变是j 行数变
		for i = startx; i < n-offset; i++ {
			matrix[i][j] = count
			count++
		}
		// 行数不变 i 列数变j--
		
	}
}
