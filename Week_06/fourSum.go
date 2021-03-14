package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	numsLen := len(nums)
	if numsLen < 4 {
		return nil
	}
	var res [][]int
	sort.Ints(nums) //排序
	for a := 0; a < numsLen-3; a++ {
		if a > 0 && nums[a] == nums[a-1] { //a去重
			continue
		}
		for b := a + 1; b < numsLen-2; b++ {
			if b > a+1 && nums[b] == nums[b-1] { //b去重
				continue
			}
			c := b + 1
			d := numsLen - 1
			for c < d {
				if c > b+1 && c < numsLen && nums[c] == nums[c-1] { //c去重
					c++
					continue
				}
				if d < numsLen-1 && d >= c && nums[d] == nums[d+1] { //d去重
					d--
					continue
				}
				sum := nums[a] + nums[b] + nums[c] + nums[d]
				//fmt.Println("a=%d,b=%d,c=%d,d=%d,sum=%d", nums[a], nums[b], nums[c], nums[d], sum)
				if sum == target {
					//fmt.Println("a=%d,b=%d,c=%d,d=%d", a, b, c, d)
					res = append(res, []int{nums[a], nums[b], nums[c], nums[d]})
					c++
					d--
				} else if sum > target {
					d--
				} else {
					c++
				}
			}
		}
	}
	return res
}

func main() {
	nums := []int{-2, -1, -1, 1, 1, 2, 2}
	res := fourSum(nums, 0)
	fmt.Println(res)
}
