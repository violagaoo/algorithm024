package main

//120. 三角形最小路径和
func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1]) //dp方程
		}
	}
	return triangle[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
