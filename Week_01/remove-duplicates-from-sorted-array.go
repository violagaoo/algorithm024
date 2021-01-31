//26. 删除排序数组中的重复项

func removeDuplicates(nums []int) int {
	var j = 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1
}




