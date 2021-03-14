package main

//213. 打家劫舍 II
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return max(myRob(nums[1:]), myRob(nums[:len(nums)-1]))
}

func myRob(nums []int) int {
	cur, pre := 0, 0
	for _, v := range nums {
		cur, pre = max(pre+v, cur), cur
	}
	return cur
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
