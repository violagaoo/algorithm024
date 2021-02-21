package main

import "fmt"

//所有子集
func subsets(nums []int) [][]int {
	res := [][]int{}

	var dfs func(i int, list []int)
	dfs = func(i int, list []int) {
		if i == len(nums) {
			tmp := make([]int, len(list))
			copy(tmp, list)
			res = append(res, tmp) //因为时
			return
		}
		list = append(list, nums[i])
		dfs(i+1, list)
		list = list[:len(list)-1]
		dfs(i+1, list)
	}
	dfs(0, []int{})

	return res
}

func main() {
	nums := []int{1, 2, 3}
	ret := subsets(nums)
	fmt.Println(ret) //[[1 2 3] [1 2] [1 3] [1] [2 3] [2] [3] []]
}
