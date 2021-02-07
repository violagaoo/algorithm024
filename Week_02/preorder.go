package main

import "fmt"

/**
 * Definition for a Node.
 */
type Node struct {
	Val      int
	Children []*Node
}

//递归解法
// var ret []int

// func preorder(root *Node) []int {
// 	ret = []int{}
// 	dfs(root)
// 	return ret
// }

// func dfs(root *Node) {
// 	if root == nil {
// 		return
// 	}
// 	ret = append(ret, root.Val)
// 	for _, n := range root.Children {
// 		dfs(n)
// 	}
// }

//非递归写法
func preorder(root *Node) []int {
	ret := []int{}
	stack := []*Node{root}
	for 0 < len(stack) {
		for root != nil {
			ret = append(ret, root.Val) //前序输出
			if 0 == len(root.Children) {
				break
			}
			for i := len(root.Children) - 1; 0 < i; i-- {
				stack = append(stack, root.Children[i]) //倒着放
			}
			root = root.Children[0]
		}
		root = stack[len(stack)-1]   //stack中最后一个放入的
		stack = stack[:len(stack)-1] //去掉stack最后一个
	}
	return ret
}

func main4() {
	nums := []int{3, 2, 4}
	nums = append(nums, 7)
	nums = append(nums, 8)
	nums = append(nums, 9)
	//fmt.Println(nums[len(nums)-1])输出9
	fmt.Println(nums[:len(nums)-1]) //输出[3 2 4 7 8], 左闭右开
}
