package main

//import strings
//括号生成
//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

func generateParenthesis(n int) []string {
	var generate func(left, right, n int, s string)
	ret := make([]string, 0)
	generate = func(left, right, n int, s string) {
		if left == n && right == n {
			ret = append(ret, s)
		}
		if left < n {
			generate(left+1, right, n, s+"(")
		}
		if left > right {
			generate(left, right+1, n, s+")")
		}
	}
	generate(0, 0, n, "")
	return ret
}

// func mian6() {

// }
