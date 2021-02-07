package main

import "fmt"

//twoSum

func twoSum(nums []int, target int) []int {
	ret := make([]int, 0)
	if len(nums) < 2 {
		return ret
	}
	mapTmp := make(map[int]int, len(nums))
	for i, v := range nums {
		mapTmp[v] = i
	}
	for i := 0; i < len(nums); i++ {
		index, ok := mapTmp[target-nums[i]]
		if ok && i != index {
			ret = append(ret, i)
			ret = append(ret, index)
			return ret
		}
	}
	return ret
}

func main() {
	nums := []int{3, 2, 4}
	ret := twoSum(nums, 6)
	fmt.Println(ret)
}
