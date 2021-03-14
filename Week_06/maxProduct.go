package main

//152. 乘积最大子数组
//https://leetcode-cn.com/problems/maximum-product-subarray/

func maxProduct(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	imax, imin, ret := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 && imax != imin {
			tmp := imax
			imax = imin
			imin = tmp
		}
		imax = max(imax*nums[i], nums[i])
		imin = min(imin*nums[i], nums[i])
		ret = max(ret, imax)
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
