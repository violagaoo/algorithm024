package main

//1143. 最长公共子序列
//字符串的动态规划，转化成二维数组
func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}

	dp := make([]int, len(text2)+1)

	for i := 1; i <= len(text1); i++ {
		last := 0
		for j := 1; j <= len(text2); j++ {
			tmp := dp[j] // tmp 记录的是二维dp矩阵正上方的值
			if text1[i-1] == text2[j-1] {
				dp[j] = last + 1 // last 记录的是二维dp矩阵左上方的值
			} else {
				dp[j] = max(tmp, dp[j-1])
			}
			last = tmp
		}
	}

	return dp[len(text2)]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
