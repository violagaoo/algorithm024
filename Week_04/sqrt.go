package main

//实现 int sqrt(int x) 函数。
//计算并返回x的平方根，其中x是非负整数。
//由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

//方法1，二分法
func mySqrt1(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	left := int64(1)
	right := int64(x)
	mid := left
	for left <= right {
		mid = left + (right-left)/2
		if mid*mid > int64(x) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return int(right)
}

//方法二，牛顿迭代法
func mySqrt2(x int) int {
	tmp := int64(x)
	for tmp*tmp > int64(x) {
		tmp = (tmp + int64(x)/tmp) / 2
	}
	return int(tmp)
}
