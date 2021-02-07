package main

//二叉树的中序遍历
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ret []int

func inorderTraversal(root *TreeNode) []int {
	ret = []int{}
	dfs(root)
	return ret
}

func dfs(root *TreeNode) {
	if root != nil {
		dfs(root.Left)
		ret = append(ret, root.Val)
		dfs(root.Right)
	}
}

// func main() {

// }
