package main

//53. 最大子序和
func maxSubArray(nums []int) int {
	ret := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + max(0, nums[i-1]) //dp方程
		if nums[i] > ret {
			ret = nums[i]
		}
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
