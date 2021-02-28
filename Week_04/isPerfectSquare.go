package main

//leetcode 367. 有效的完全平方数
//牛顿迭代法,类似sqrt.go
func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}
	x := int64(num) / 2
	for x*x > int64(num) {
		x = (x + int64(num)/x) / 2
	}
	return x*x == int64(num)
}

// public boolean isPerfectSquare(int num) {
//   if (num < 2) return true;

//   long x = num / 2;
//   while (x * x > num) {
// 	x = (x + num / x) / 2;
//   }
//   return (x * x == num);
// }
