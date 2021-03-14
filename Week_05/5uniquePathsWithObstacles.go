package main

//63. 不同路径 II,类似leetcode62,就是加了障碍物

//62.一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为 “Start” ）。
//机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为 “Finish” ）。
//问总共有多少条不同的路径

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m < 1 {
		return 0
	}
	n := len(obstacleGrid[0])
	tmp := make([]int, n+1) //防止下面j-1越界，
	tmp[1] = 1
	for i := 0; i < m; i++ {
		for j := 1; j <= n; j++ {
			if obstacleGrid[i][j-1] == 1 {
				tmp[j] = 0
			} else {
				tmp[j] += tmp[j-1]
			}
		}
	}
	return tmp[n]
}
