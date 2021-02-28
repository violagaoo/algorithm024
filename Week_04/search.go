package main

//leetcode33：搜索旋转排序数组
func search(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}
	low := 0
	high := len(nums) - 1
	mid := low
	for low <= high {//注意等号
		mid = low + (high-low)/2
		if nums[mid] == target {
			return mid
		}
		if nums[low] <= nums[mid] { //左边有序，左边查找
			if target >= nums[low] && target < nums[mid] { //注意等号，等号加在不是mid上，mid上面判断过了
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else { //如果右边有序，在右边查找
			if target <= nums[high] && target > nums[mid] { //注意等号
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}
[4,5,6,7,0,1,2]
0