package main

//leetcode62. 不同路径
//一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
//问总共有多少条不同的路径

func uniquePaths(m int, n int) int {
	newRow := make([]int, n) //存储当前行
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				newRow[j] = 1
			} else {
				newRow[j] += newRow[j-1] //上一行的值+左边的值
			}
		}
	}
	return newRow[n-1]
}
